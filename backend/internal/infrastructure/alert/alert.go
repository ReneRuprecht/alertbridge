package alert

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/application"
)

func FindAlertsByInstance(uc *application.FindAlertsByInstanceUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		instance := r.URL.Query().Get("instance")
		if instance == "" {
			http.Error(w, "instance is required", http.StatusBadRequest)
		}

		ctx := r.Context()
		alerts, err := uc.Execute(ctx, instance)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dto := FindAlertsByInstanceDto{
			Instance: instance,
			Alerts:   alerts,
		}

		json.NewEncoder(w).Encode(dto)
	}
}
