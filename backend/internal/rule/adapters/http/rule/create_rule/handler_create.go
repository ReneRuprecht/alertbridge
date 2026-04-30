package rule

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

func HandleCreateRule(uc application.CreateRuleUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

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
