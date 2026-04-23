package rule

import (
	"github.com/google/uuid"
	domain "github.com/reneruprecht/alertbridge/backend/internal/domain/rule"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

func toDomain(rule postgres_db.Rule) (domain.Rule, error) {

	name, err := domain.NewRuleName(rule.Name)

	if err != nil {
		return domain.Rule{}, err
	}

	priority, err := domain.NewRulePriority(int(rule.Priority))

	if err != nil {
		return domain.Rule{}, err
	}

	return domain.Rule{
		ID:          domain.RuleId(rule.ID),
		Name:        name,
		Description: rule.Description,
		Priority:    priority,
		Enabled:     rule.Enabled,
	}, nil

}

func toDbRule(r domain.Rule) postgres_db.Rule {
	return postgres_db.Rule{
		ID:          uuid.UUID(r.ID),
		Name:        string(r.Name),
		Description: r.Description,
		Priority:    int32(r.Priority),
		Enabled:     r.Enabled,
	}
}
