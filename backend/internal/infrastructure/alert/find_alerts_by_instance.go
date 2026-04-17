package alert

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/application"
)

func HandleFindAlertsByInstance(uc application.FindAlertsByInstanceUseCaseInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

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

		dto := toFindAlertsByInstanceDto(alerts, instance)
		json.NewEncoder(w).Encode(dto)
	}
}
