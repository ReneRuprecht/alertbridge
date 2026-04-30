package domain

import (
	"fmt"
)

type ActionType string

const (
	ActionTypeWebhook ActionType = "webhook"
)

func errorActionTypeInvalid(t ActionType) error {
	return fmt.Errorf("invalid action type: %s", t)
}

func NewActionType(action string) (ActionType, error) {
	actionType := ActionType(action)

	switch actionType {
	case ActionTypeWebhook:
		return actionType, nil
	default:
		return "", errorActionTypeInvalid(actionType)
	}
}
