package alertmanager

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/application"
)

func HandleWebhook(uc *application.ReceiveAlertUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		alerts, err := toDomain(req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := uc.Execute(alerts); err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
