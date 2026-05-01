package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type RuleRepository struct {
	queries *postgres_db.Queries
}

func NewRuleRepository(queries *postgres_db.Queries) *RuleRepository {
	return &RuleRepository{queries: queries}
}

func (r *RuleRepository) Save(ctx context.Context, rule domain.Rule) error {

	ruleEntity := toRuleRepositoryEntity(rule)

	return r.queries.InsertRule(ctx, postgres_db.InsertRuleParams{
		ID:          ruleEntity.ID,
		Name:        ruleEntity.Name,
		Description: ruleEntity.Description,
		Priority:    ruleEntity.Priority,
		Enabled:     ruleEntity.Enabled,
	})
}

func (r *RuleRepository) FindByID(ctx context.Context, ruleID domain.RuleId) (domain.Rule, error) {

	row, err := r.queries.FindRuleByID(ctx, uuid.UUID(ruleID))

	if err == sql.ErrNoRows {
		return domain.Rule{}, application.ErrorRuleNotFound
	}
	if err != nil {
		return domain.Rule{}, err
	}
	return domain.Rule{
		ID:          domain.RuleId(row.ID),
		Name:        domain.RuleName(row.Name),
		Description: row.Description,
		Priority:    domain.RulePriority(row.Priority),
		Enabled:     row.Enabled,
	}, nil

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
