package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

type SaveActionUseCase interface {
	Execute(ctx context.Context, action domain.Action) error
}

type saveActionUseCase struct {
	repo ActionRepositoryWriter
}

func NewSaveActionUseCase(repo ActionRepositoryWriter) *saveActionUseCase {
	return &saveActionUseCase{repo: repo}
}

func (s *saveActionUseCase) Execute(ctx context.Context, action domain.Action) error {
	return s.repo.Save(ctx, action)
}
