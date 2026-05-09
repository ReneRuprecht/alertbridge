package domain

import "github.com/google/uuid"

type ConditionID uuid.UUID

func NewConditionID() (ConditionID, error) {

	id, err := uuid.NewV7()

	if err != nil {
		return ConditionID{}, err
	}

	return ConditionID(id), nil
}

func (c ConditionID) String() string {
	return uuid.UUID(c).String()
}
