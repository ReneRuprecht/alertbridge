package alert

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/application"
)

func HandleListActiveAlerts(uc *application.ListActiveAlertsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()
		alerts, err := uc.Execute(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dto := ListActiveAlertsDto{
			Alerts: alerts,
		}

		json.NewEncoder(w).Encode(dto)
	}
}
