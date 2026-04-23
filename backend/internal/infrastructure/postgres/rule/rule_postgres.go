package rule

import (
	"context"

	domain "github.com/reneruprecht/alertbridge/backend/internal/domain/rule"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

type RuleRepository struct {
	queries *postgres_db.Queries
}

func NewRuleRepository(queries *postgres_db.Queries) *RuleRepository {
	return &RuleRepository{queries: queries}
}

func (r *RuleRepository) Save(ctx context.Context, rule domain.Rule) error {

	ruleDto := toDbRule(rule)

	return r.queries.InsertRule(ctx, postgres_db.InsertRuleParams{
		ID:          ruleDto.ID,
		Name:        ruleDto.Name,
		Description: ruleDto.Description,
		Priority:    ruleDto.Priority,
		Enabled:     ruleDto.Enabled,
	})
}

func (r *RuleRepository) List(ctx context.Context) ([]domain.Rule, error) {

	rows, err := r.queries.ListRules(ctx)

	if err != nil {
		return nil, err
	}

	rules := make([]domain.Rule, len(rows))

	for i, row := range rows {
		rule, err := toDomain(row)

		if err != nil {
			return nil, err
		}

		rules[i] = rule

	}

	return rules, nil
}
