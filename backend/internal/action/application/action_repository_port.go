package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

type ActionRepositoryWriter interface {
	Save(ctx context.Context, action domain.Action) error
}

type ActionRepositoryReader interface {
	List(ctx context.Context) ([]domain.Action, error)
}

type ActionRepositoryFinder interface {
	FindByID(ctx context.Context, actionID domain.ActionID) (domain.Action, error)
}
