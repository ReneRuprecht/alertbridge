package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type ListAlertsByInstanceUseCase interface {
	Execute(ctx context.Context, instance string) ([]domain.Alert, error)
}

type listAlertsByInstanceUseCase struct {
	repo AlertRepositoryReader
}

func NewListAlertsByInstanceUseCase(repo AlertRepositoryReader) *listAlertsByInstanceUseCase {
	return &listAlertsByInstanceUseCase{repo: repo}
}

func (uc *listAlertsByInstanceUseCase) Execute(ctx context.Context, instance string) ([]domain.Alert, error) {
	return uc.repo.ListByInstance(ctx, instance)
}
