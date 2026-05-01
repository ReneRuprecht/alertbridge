package application

import (
	"context"
	"errors"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

var ErrorConditionsNotFound = errors.New("conditions not found")

type RuleConditionRepository interface {
	Save(context context.Context, condition domain.Condition) error
	List(context context.Context) ([]domain.Condition, error)
	ListByRuleID(context context.Context, ruleID domain.RuleId) ([]domain.Condition, error)
}
