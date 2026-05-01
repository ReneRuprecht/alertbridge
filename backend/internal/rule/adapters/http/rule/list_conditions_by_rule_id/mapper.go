package listconditionsbyruleid

import (
	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

func toRuleID(ruleID string) (domain.RuleId, error) {
	parsedID, err := uuid.Parse(ruleID)

	if err != nil {
		return domain.RuleId{}, err
	}

	return domain.RuleId(parsedID), nil
}

func toListConditionsByRuleIDResponse(dto application.ListConditionsByRuleIDUseCaseDto) ListConditionsByRuleIDResponse {

	rule := Rule{ID: uuid.UUID(dto.Rule.ID).String(), Name: dto.Rule.Name.String(), Description: dto.Rule.Description, Priority: dto.Rule.Priority.Int(), Enabled: dto.Rule.Enabled}

	conditions := make([]Condition, len(dto.Conditions))

	for i, c := range dto.Conditions {
		condition := Condition{ID: uuid.UUID(c.ID).String(), Name: string(c.Name), Operator: string(c.Operator), Field: string(c.Field), Value: string(c.Value)}
		conditions[i] = condition
	}

	return ListConditionsByRuleIDResponse{Rule: rule, Conditions: conditions}

}
