package listruleconditions

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
)

func HandleListRuleConditions(uc application.ListRuleConditionsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ctx := r.Context()
		conditions, err := uc.Execute(ctx)

		if err != nil {
			switch {
			case errors.Is(err, application.ErrorConditionsNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		response := toListRuleConditionsResponse(conditions)
		json.NewEncoder(w).Encode(response)
	}
}
