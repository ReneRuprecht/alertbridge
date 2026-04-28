package rule

import "github.com/reneruprecht/alertbridge/backend/internal/rule/domain"

func toListRulesResponse(rules []domain.Rule) ListRulesResponse {

	response := ListRulesResponse{Rules: []Rule{}}

	for _, rule := range rules {

		r := Rule{
			ID:          rule.ID.String(),
			Name:        string(rule.Name),
			Description: rule.Description,
			Priority:    rule.Priority.Int(),
			Enabled:     rule.Enabled,
		}

		response.Rules = append(response.Rules, r)
	}

	return response

}
