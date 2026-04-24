package rule

import "github.com/reneruprecht/alertbridge/backend/internal/rule/domain"

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
