package rule

import (
	domain "github.com/reneruprecht/alertbridge/backend/internal/domain/rule"
)

func toDomain(request CreateRuleRequest) (domain.Rule, error) {

	id := domain.NewRuleId()
	name, err := domain.NewRuleName(request.Name)

	if err != nil {
		return domain.Rule{}, err
	}

	priority, err := domain.NewRulePriority(request.Priority)

	if err != nil {
		return domain.Rule{}, err
	}

	return domain.Rule{
		ID:          id,
		Name:        name,
		Description: request.Description,
		Priority:    priority,
		Enabled:     request.Enabled,
	}, nil
}

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
