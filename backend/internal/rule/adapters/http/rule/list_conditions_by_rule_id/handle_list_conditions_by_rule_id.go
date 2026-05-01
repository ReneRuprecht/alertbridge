package listconditionsbyruleid

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
)

func HandleListConditionsByRuleID(uc application.ListConditionsByRuleIDUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ruleIDValue := r.PathValue("id")
		if ruleIDValue == "" {
			http.Error(w, "rule id is required", http.StatusBadRequest)
		}

		ctx := r.Context()
		ruleID, err := toRuleID(ruleIDValue)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		dto, err := uc.Execute(ctx, ruleID)
		if err != nil {
			switch {
			case errors.Is(err, application.ErrorRuleNotFound),
				errors.Is(err, application.ErrorConditionsNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			default:

				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		response := toListConditionsByRuleIDResponse(dto)

		json.NewEncoder(w).Encode(response)
	}
}
