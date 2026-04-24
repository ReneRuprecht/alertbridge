package application

import (
	"context"
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type CreateRuleUseCaseInterface interface {
	Execute(ctx context.Context, rule domain.Rule) error
}

type CreateRuleUseCase struct {
	repo RuleRepository
}

func NewCreateRuleUseCase(repo RuleRepository) *CreateRuleUseCase {
	return &CreateRuleUseCase{repo: repo}
}

func (uc *CreateRuleUseCase) Execute(ctx context.Context, rule domain.Rule) error {

	err := uc.repo.Save(ctx, rule)

	if err != nil {
		log.Printf("CreateRuleUseCase error %v", err)
		return err
	}
	return nil
}
