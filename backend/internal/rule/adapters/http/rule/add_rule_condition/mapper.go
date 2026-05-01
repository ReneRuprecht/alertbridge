package addrulecondition

import (
	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

func toDomain(request AddRuleConditionRequest) (domain.Condition, error) {
	ruleUUID, err := uuid.Parse(request.RuleID)

	if err != nil {
		return domain.Condition{}, err
	}
	ruleID := domain.RuleId(ruleUUID)

	id, err := domain.NewConditionID()
	if err != nil {
		return domain.Condition{}, err
	}
	name, err := domain.NewConditionName(request.Name)
	if err != nil {
		return domain.Condition{}, err
	}
	operator, err := domain.NewConditionOperator(request.Operator)
	if err != nil {
		return domain.Condition{}, err
	}
	field, err := domain.NewConditionField(request.Field)
	if err != nil {
		return domain.Condition{}, err
	}
	value, err := domain.NewConditionValue(request.Value)
	if err != nil {
		return domain.Condition{}, err
	}
	return domain.Condition{
		ID:       id,
		RuleID:   ruleID,
		Name:     name,
		Operator: operator,
		Field:    field,
		Value:    value,
	}, nil

}
