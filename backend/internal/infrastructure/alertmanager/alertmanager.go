package alertmanager

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req WebhookRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	log.Printf("Received %d alerts", len(req.Alerts))

	w.WriteHeader(http.StatusOK)
}
