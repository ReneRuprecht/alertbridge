package rule

import "github.com/reneruprecht/alertbridge/backend/internal/rule/domain"

func toListRulesDto(rules []domain.Rule) ListRulesResponse {

	dto := ListRulesResponse{Rules: []Rule{}}

	for _, rule := range rules {

		r := Rule{
			ID:          rule.ID.String(),
			Name:        string(rule.Name),
			Description: rule.Description,
			Priority:    rule.Priority.Int(),
			Enabled:     rule.Enabled,
		}

		dto.Rules = append(dto.Rules, r)
	}

	return dto

}
