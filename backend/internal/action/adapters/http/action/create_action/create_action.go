package createaction

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/action/application"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

func HandleCreateAction(uc application.SaveActionUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ctx := r.Context()

		var req createActionRequest
		json.NewDecoder(r.Body).Decode(&req)

		action, err := toDomain(req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := uc.Execute(ctx, action); err != nil {
			switch {
			case errors.Is(err, domain.ErrorActionConfigEmpty),
				errors.Is(err, domain.ErrorActionNameEmpty):
				http.Error(w, "action validation failed", http.StatusBadRequest)

			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)

			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
