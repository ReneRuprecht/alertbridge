package main

import (
	"log"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alertmanager"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/alerts", alertmanager.HandleWebhook)

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
