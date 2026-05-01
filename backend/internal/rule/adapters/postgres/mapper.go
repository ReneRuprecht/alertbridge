package postgres

import (
	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
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

func toRuleRepositoryEntity(r domain.Rule) postgres_db.Rule {
	return postgres_db.Rule{
		ID:          uuid.UUID(r.ID),
		Name:        r.Name.String(),
		Description: r.Description,
		Priority:    int32(r.Priority.Int()),
		Enabled:     r.Enabled,
	}
}

func toRuleConditionRepositoryEntity(condition domain.Condition) postgres_db.RuleCondition {
	return postgres_db.RuleCondition{
		ID:       uuid.UUID(condition.ID),
		RuleID:   uuid.UUID(condition.RuleID),
		Name:     string(condition.Name),
		Operator: string(condition.Operator),
		Field:    string(condition.Field),
		Value:    string(condition.Value),
	}
}

func toConditionDomain(row postgres_db.RuleCondition) (domain.Condition, error) {

	parsedConditionID, err := uuid.Parse(row.ID.String())

	if err != nil {
		return domain.Condition{}, err
	}
	parsedRuleID, err := uuid.Parse(row.RuleID.String())
	if err != nil {
		return domain.Condition{}, err
	}

	name, err := domain.NewConditionName(row.Name)
	if err != nil {
		return domain.Condition{}, err
	}

	operator, err := domain.NewConditionOperator(row.Operator)
	if err != nil {
		return domain.Condition{}, err
	}

	field, err := domain.NewConditionField(row.Field)
	if err != nil {
		return domain.Condition{}, err
	}

	value, err := domain.NewConditionValue(row.Value)
	if err != nil {
		return domain.Condition{}, err
	}

	return domain.Condition{
		ID:       domain.ConditionID(parsedConditionID),
		RuleID:   domain.RuleId(parsedRuleID),
		Name:     name,
		Operator: operator,
		Field:    field,
		Value:    value,
	}, nil

}
