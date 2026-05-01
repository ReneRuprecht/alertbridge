package listruleconditions

import (
	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

func toListRuleConditionsResponse(conditions []domain.Condition) ListRuleConditionsResponse {

	response := ListRuleConditionsResponse{Conditions: []Condition{}}

	for _, c := range conditions {

		condition := Condition{
			ID:       uuid.UUID(c.ID).String(),
			RuleID:   c.RuleID.String(),
			Name:     string(c.Name),
			Operator: string(c.Operator),
			Field:    string(c.Field),
			Value:    string(c.Value),
		}

		response.Conditions = append(response.Conditions, condition)
	}

	return response

}
