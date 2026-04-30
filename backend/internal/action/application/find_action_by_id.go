package application

import (
	"context"
	"errors"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)


var ErrorActionNotFound = errors.New("action does not exists")

type FindActionByIDUseCase interface {
	Execute(ctx context.Context, actionID domain.ActionID) (domain.Action, error)
}

type findActionByIDUseCase struct {
	repo ActionRepositoryFinder
}

func NewFindActionByIDUseCase(repo ActionRepositoryFinder) *findActionByIDUseCase {
	return &findActionByIDUseCase{repo: repo}
}

func (l *findActionByIDUseCase) Execute(ctx context.Context, actionID domain.ActionID) (domain.Action, error) {
	return l.repo.FindByID(ctx, actionID)

}
