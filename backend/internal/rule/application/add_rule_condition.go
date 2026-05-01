package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type AddRuleConditionUseCase interface {
	Execute(ctx context.Context, condition domain.Condition) error
}

type addRuleConditionUseCase struct {
	ruleFinder RuleFinder
	repo       RuleConditionRepository
}

func NewAddRuleConditionUseCase(ruleFinder RuleFinder, repo RuleConditionRepository) *addRuleConditionUseCase {
	return &addRuleConditionUseCase{ruleFinder: ruleFinder, repo: repo}
}

func (uc *addRuleConditionUseCase) Execute(ctx context.Context, condition domain.Condition) error {

	if _, err := uc.ruleFinder.FindByID(ctx, condition.RuleID); err != nil {
		return err
	}

	return uc.repo.Save(ctx, condition)
}
