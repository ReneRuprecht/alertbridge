package rule

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
)

func HandleListRules(uc application.ListRulesUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ctx := r.Context()
		rules, err := uc.Execute(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := toListRulesResponse(rules)
		json.NewEncoder(w).Encode(response)
	}
}
