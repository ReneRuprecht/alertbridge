package postgres

import (
	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

func toUUID(actionID domain.ActionID) uuid.UUID {
	return uuid.UUID(actionID)
}

func toActionRepositoryEntity(action domain.Action) actionRepositoryEntity {
	id := toUUID(action.ID)

	return actionRepositoryEntity{
		ID:          id,
		Name:        string(action.Name),
		Description: string(action.Description),
		Type:        string(action.Type),
		Config:      action.Config,
	}
}

func toDomain(row postgres_db.Action) (domain.Action, error) {
	name, err := domain.NewActionName(row.Name)
	if err != nil {
		return domain.Action{}, err
	}

	desc := domain.NewActionDescription(row.Description)

	actionType, err := domain.NewActionType(row.Type)
	if err != nil {
		return domain.Action{}, err
	}

	return domain.Action{
		ID:          domain.ActionID(row.ID),
		Name:        name,
		Description: desc,
		Type:        actionType,
		Config:      row.Config,
	}, nil
}
