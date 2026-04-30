package findactionbyid

import (
	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

func toActionID(id string) (domain.ActionID, error) {

	parsedID, err := uuid.Parse(id)

	if err != nil {
		return domain.ActionID{}, err
	}

	return domain.ActionID(parsedID), nil

}

func toFindActionByIDResponse(action domain.Action) findActionByIDResponse {
	return findActionByIDResponse{
		ID:          uuid.UUID(action.ID).String(),
		Name:        string(action.Name),
		Description: string(action.Description),
		Type:        string(action.Type),
		Config:      action.Config,
	}
}
