package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/sashabaranov/go-openai"
)

func main() {
    router := gin.Default()
    client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

    router.POST("/chat", func(c *gin.Context) {
        var request struct {
            Message string `json:"message"`
        }

        if err := c.BindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        resp, err := client.CreateChatCompletion(
            c,
            openai.ChatCompletionRequest{
                Model: openai.GPT3Dot5Turbo,
                Messages: []openai.ChatCompletionMessage{
                    {
                        Role:    openai.ChatMessageRoleUser,
                        Content: request.Message,
                    },
                },
            },
        )

        if err != nil {
            log.Printf("ChatCompletion error: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"response": resp.Choices[0].Message.Content})
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    router.Run(":" + port)
}