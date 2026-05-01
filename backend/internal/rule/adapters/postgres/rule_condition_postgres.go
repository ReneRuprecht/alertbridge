package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type RuleConditionRepository struct {
	queries *postgres_db.Queries
}

func NewRuleConditionRepository(queries *postgres_db.Queries) *RuleConditionRepository {
	return &RuleConditionRepository{queries: queries}
}

func (r *RuleConditionRepository) Save(ctx context.Context, condition domain.Condition) error {

	ruleConditionEntity := toRuleConditionRepositoryEntity(condition)
	return r.queries.InsertRuleCondition(ctx, postgres_db.InsertRuleConditionParams(
		ruleConditionEntity,
	))

}

func (r *RuleConditionRepository) List(ctx context.Context) ([]domain.Condition, error) {

	rows, err := r.queries.ListRuleConditions(ctx)

	if err == sql.ErrNoRows {
		return []domain.Condition{}, application.ErrorConditionsNotFound
	}
	if err != nil {
		return []domain.Condition{}, err
	}

	conditions := make([]domain.Condition, len(rows))

	for i, row := range rows {
		condition, err := toConditionDomain(row)

		if err != nil {
			return []domain.Condition{}, err
		}
		conditions[i] = condition
	}
	return conditions, nil

}

func (r *RuleConditionRepository) ListByRuleID(ctx context.Context, ruleID domain.RuleId) ([]domain.Condition, error) {
	parsedRuleID, err := uuid.Parse(ruleID.String())

	if err != nil {
		return []domain.Condition{}, err
	}

	rows, err := r.queries.ListRuleConditionsByRuleID(ctx, parsedRuleID)

	if err != nil {
		return []domain.Condition{}, err
	}

	conditions := make([]domain.Condition, len(rows))

	for i, row := range rows {
		condition, err := toConditionDomain(row)

		if err != nil {
			return []domain.Condition{}, err
		}
		conditions[i] = condition
	}
	return conditions, nil

}
