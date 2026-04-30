package createaction

import (
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

func toDomain(actionRequest createActionRequest) (domain.Action, error) {
	id, err := domain.NewActionID()

	if err != nil {
		return domain.Action{}, err
	}

    log.Println(actionRequest)
	name, err := domain.NewActionName(actionRequest.Name)

	if err != nil {
		return domain.Action{}, err
	}

	description := domain.NewActionDescription(actionRequest.Description)

	actionType, err := domain.NewActionType(actionRequest.Type)
	if err != nil {
		return domain.Action{}, err
	}

	config, err := domain.NewActionConfig(actionRequest.Config)
	if err != nil {
		return domain.Action{}, err
	}

	return domain.Action{ID: id, Name: name, Description: description, Type: actionType, Config: config}, nil

}
