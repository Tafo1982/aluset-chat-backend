package main

import (
    "log"
    "net/http"
    "github.com/sashabaranov/go-openai"
)

func main() {
    client := openai.NewClient("API_KEY_TUAJ_KETU")

    http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
        // Logjika për chatbot-in
        // Për shembull, lexoni pyetjen nga kërkesa dhe përgjigju me OpenAI
        // Këtu mund ta shtosh kodin për trajtimin e kërkesave dhe përgjigjeve
    })

    log.Println("Serveri po funksionon në portin 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
