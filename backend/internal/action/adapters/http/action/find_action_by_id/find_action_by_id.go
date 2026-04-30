package findactionbyid

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/action/application"
)

func HandleFindActionByID(uc application.FindActionByIDUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
		}

		ctx := r.Context()
		domainID, err := toActionID(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		action, err := uc.Execute(ctx, domainID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := toFindActionByIDResponse(action)
		json.NewEncoder(w).Encode(response)
	}
}
