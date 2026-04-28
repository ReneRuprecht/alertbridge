package alertmanager

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
)

func HandleWebhook(uc application.SaveAlertsWithCacheUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		ctx := r.Context()
		if err := uc.Execute(ctx, alerts); err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
