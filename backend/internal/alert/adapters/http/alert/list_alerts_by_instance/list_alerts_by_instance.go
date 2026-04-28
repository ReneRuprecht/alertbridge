package alert

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
)

func HandleListAlertsByInstance(uc application.ListAlertsByInstanceUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		instance := r.PathValue("instance")
		if instance == "" {
			http.Error(w, "instance is required", http.StatusBadRequest)
		}

		ctx := r.Context()
		alerts, err := uc.Execute(ctx, instance)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := toListAlertsByInstanceResponse(alerts, instance)
		json.NewEncoder(w).Encode(response)
	}
}
