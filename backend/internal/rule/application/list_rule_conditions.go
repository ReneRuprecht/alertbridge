package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type ListRuleConditionsUseCase interface {
	Execute(ctx context.Context) ([]domain.Condition, error)
}

type listRuleConditionsUseCase struct {
	repo RuleConditionRepository
}

func NewListRuleConditionsUseCase(repo RuleConditionRepository) *listRuleConditionsUseCase {
	return &listRuleConditionsUseCase{repo: repo}
}

func (uc *listRuleConditionsUseCase) Execute(ctx context.Context) ([]domain.Condition, error) {

	return uc.repo.List(ctx)
}
