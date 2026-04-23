package rule

import (
	"encoding/json"
	"net/http"

	application "github.com/reneruprecht/alertbridge/backend/internal/application/rule"
)

func HandleListRules(uc application.ListRulesUseCaseInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		ctx := r.Context()
		rules, err := uc.Execute(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dto := toListRulesDto(rules)
		json.NewEncoder(w).Encode(dto)
	}
}
