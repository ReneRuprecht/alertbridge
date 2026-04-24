package application

import (
	"context"
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type ListRulesUseCaseInterface interface {
	Execute(ctx context.Context) ([]domain.Rule, error)
}

type ListRuleUseCase struct {
	repo RuleRepository
}

func NewListRuleUseCase(repo RuleRepository) *ListRuleUseCase {
	return &ListRuleUseCase{repo: repo}
}

func (uc *ListRuleUseCase) Execute(ctx context.Context) ([]domain.Rule, error) {

	rules, err := uc.repo.List(ctx)

	if err != nil {
		log.Printf("ListRuleUseCase error %v", err)
		return []domain.Rule{}, err
	}
	return rules, nil
}
