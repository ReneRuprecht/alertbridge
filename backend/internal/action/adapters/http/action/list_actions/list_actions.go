package listaction

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/action/application"
)

func HandleListActions(uc application.ListActionsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ctx := r.Context()
		actions, err := uc.Execute(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := toListActionsResponse(actions)
		json.NewEncoder(w).Encode(response)
	}
}
