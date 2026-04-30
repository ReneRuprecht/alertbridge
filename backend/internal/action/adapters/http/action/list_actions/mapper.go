package listaction

import (
	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

func toListActionsResponse(actions []domain.Action) listActionsResponse {

	response := listActionsResponse{}

	response.Actions = make([]action, len(actions))

	for i, v := range actions {

		response.Actions[i] = action{ID: uuid.UUID(v.ID).String(), Name: string(v.Name), Description: string(v.Description), Type: string(v.Type), Config: v.Config}

	}

	return response

}
