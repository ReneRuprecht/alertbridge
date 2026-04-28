package alert

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
)

func HandleListActiveAlerts(uc application.ListActiveAlertsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ctx := r.Context()
		alerts, err := uc.Execute(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dto := toListActiveAlertDto(alerts)
		json.NewEncoder(w).Encode(dto)
	}
}
