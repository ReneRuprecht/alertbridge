package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type ListRulesUseCase interface {
	Execute(ctx context.Context) ([]domain.Rule, error)
}

type listRuleUseCase struct {
	repo RuleRepository
}

func NewListRuleUseCase(repo RuleRepository) *listRuleUseCase {
	return &listRuleUseCase{repo: repo}
}

func (uc *listRuleUseCase) Execute(ctx context.Context) ([]domain.Rule, error) {
	return uc.repo.List(ctx)
}
