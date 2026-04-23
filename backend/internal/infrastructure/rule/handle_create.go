package rule

import (
	"encoding/json"
	"errors"
	"net/http"

	application "github.com/reneruprecht/alertbridge/backend/internal/application/rule"
	domain "github.com/reneruprecht/alertbridge/backend/internal/domain/rule"
)

func HandleCreateRule(uc application.CreateRuleUseCaseInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ctx := r.Context()

		var req CreateRuleRequest
		json.NewDecoder(r.Body).Decode(&req)

		rule, err := toDomain(req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := uc.Execute(ctx, rule); err != nil {
			switch {
			case errors.Is(err, domain.ErrorRuleNameEmpty),
				errors.Is(err, domain.ErrorRulePriorityNegative):
				http.Error(w, "role validation failed", http.StatusBadRequest)

			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)

			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
