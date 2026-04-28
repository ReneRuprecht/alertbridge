package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type CreateRuleUseCase interface {
	Execute(ctx context.Context, rule domain.Rule) error
}

type createRuleUseCase struct {
	repo RuleRepository
}

func NewCreateRuleUseCase(repo RuleRepository) *createRuleUseCase {
	return &createRuleUseCase{repo: repo}
}

func (uc *createRuleUseCase) Execute(ctx context.Context, rule domain.Rule) error {
	return uc.repo.Save(ctx, rule)
}
