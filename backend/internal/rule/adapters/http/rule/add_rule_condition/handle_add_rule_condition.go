package addrulecondition

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

func HandleAddRuleCondition(uc application.AddRuleConditionUseCase) http.HandlerFunc {
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

		var req AddRuleConditionRequest
		json.NewDecoder(r.Body).Decode(&req)

		ruleCondition, err := toDomain(req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := uc.Execute(ctx, ruleCondition); err != nil {
			switch {
			case errors.Is(err, domain.ErrorConditionNameEmpty),
				errors.Is(err, domain.ErrorConditionFieldEmpty),
				errors.Is(err, domain.ErrorConditionOperatorEmpty),
				errors.Is(err, domain.ErrorConditionOperatorInvalid):
				http.Error(w, "condition validation failed", http.StatusBadRequest)
			case errors.Is(err, application.ErrorRuleNotFound):
				http.Error(w, "rule not found", http.StatusNotFound)

			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)

			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
