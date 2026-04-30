package postgres

import (
	"context"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

type ActionRepository struct {
	queries *postgres_db.Queries
}

func NewActionRepository(queries *postgres_db.Queries) *ActionRepository {
	return &ActionRepository{queries: queries}
}

func (r *ActionRepository) Save(ctx context.Context, action domain.Action) error {

	actionRepositoryEntity := toActionRepositoryEntity(action)

	return r.queries.InsertAction(ctx, postgres_db.InsertActionParams{
		ID:          actionRepositoryEntity.ID,
		Name:        actionRepositoryEntity.Name,
		Description: actionRepositoryEntity.Description,
		Type:        actionRepositoryEntity.Type,
		Config:      actionRepositoryEntity.Config,
	})
}

func (r *ActionRepository) FindByID(ctx context.Context, actionID domain.ActionID) (domain.Action, error) {

	id := toUUID(actionID)

	row, err := r.queries.FindActionById(ctx, id)

	if err != nil {
		return domain.Action{}, err
	}

	return toDomain(row)

}

func (r *ActionRepository) List(ctx context.Context) ([]domain.Action, error) {

	rows, err := r.queries.ListActions(ctx)

	if err != nil {
		return []domain.Action{}, err
	}

	actions := make([]domain.Action, len(rows))
	for i, row := range rows {

		action, err := toDomain(row)

		if err != nil {
			return []domain.Action{}, err
		}
		actions[i] = action
	}

	return actions, nil
}
