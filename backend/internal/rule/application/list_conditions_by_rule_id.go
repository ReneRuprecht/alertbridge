package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type ListConditionsByRuleIDUseCaseDto struct {
	Rule       domain.Rule
	Conditions []domain.Condition
}

type ListConditionsByRuleIDUseCase interface {
	Execute(ctx context.Context, ruleID domain.RuleId) (ListConditionsByRuleIDUseCaseDto, error)
}

type listConditionsByRuleIDUseCase struct {
	ruleFinder    RuleFinder
	conditionRepo RuleConditionRepository
}

func NewListConditionsByRuleIDUseCase(ruleFinder RuleFinder, conditionRepo RuleConditionRepository) *listConditionsByRuleIDUseCase {
	return &listConditionsByRuleIDUseCase{ruleFinder: ruleFinder, conditionRepo: conditionRepo}
}

func (uc *listConditionsByRuleIDUseCase) Execute(ctx context.Context, ruleID domain.RuleId) (ListConditionsByRuleIDUseCaseDto, error) {

	rule, err := uc.ruleFinder.FindByID(ctx, ruleID)

	if err != nil {
		return ListConditionsByRuleIDUseCaseDto{}, err
	}
	conditions, err := uc.conditionRepo.ListByRuleID(ctx, ruleID)
	if err != nil {
		return ListConditionsByRuleIDUseCaseDto{}, err
	}

	return ListConditionsByRuleIDUseCaseDto{Rule: rule, Conditions: conditions}, nil
}
