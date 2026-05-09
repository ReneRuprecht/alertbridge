package application

import (
	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type FieldResolver func(alert application.AlertCacheDto) string

var fieldResolvers = map[string]FieldResolver{
	"status":   func(a application.AlertCacheDto) string { return string(a.Status) },
	"severity": func(a application.AlertCacheDto) string { return string(a.Severity) },
}

type MatchRuleConditionUseCase interface {
	Execute(alertCacheDto application.AlertCacheDto, condition domain.Condition) bool
}

type matchRuleConditionUseCase struct {
}

func NewMatchRuleConditionUseCase() *matchRuleConditionUseCase {
	return &matchRuleConditionUseCase{}
}

func (uc *matchRuleConditionUseCase) Execute(alertCacheDto application.AlertCacheDto, condition domain.Condition) bool {

	resolver, ok := fieldResolvers[string(condition.Field)]
	if !ok {
		return false
	}

	alertValue := resolver(alertCacheDto)

	return condition.Operator.Apply(alertValue, string(condition.Value))

}
