package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

type ListActionsUseCase interface {
	Execute(ctx context.Context) ([]domain.Action, error)
}

type listActionsUseCase struct {
	repo ActionRepositoryReader
}

func NewListActionsUseCase(repo ActionRepositoryReader) *listActionsUseCase {
	return &listActionsUseCase{repo: repo}
}

func (l *listActionsUseCase) Execute(ctx context.Context) ([]domain.Action, error) {
	return l.repo.List(ctx)
}
