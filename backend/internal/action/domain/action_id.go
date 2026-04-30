package domain

import "github.com/google/uuid"

type ActionID uuid.UUID

func NewActionID() (ActionID, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return ActionID{}, err
	}

	return ActionID(id), nil
}
