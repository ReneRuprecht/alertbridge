package domain

import "github.com/google/uuid"

type RuleId uuid.UUID

func NewRuleId() (RuleId, error) {

	id, err := uuid.NewV7()

	if err != nil {
		return RuleId{}, err
	}

	return RuleId(id), nil
}
func (r RuleId) String() string {
	return uuid.UUID(r).String()
}
