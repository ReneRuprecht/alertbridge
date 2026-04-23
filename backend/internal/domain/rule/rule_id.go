package rule

import "github.com/google/uuid"

type RuleId uuid.UUID

func NewRuleId() RuleId {

	id, _ := uuid.NewV7()

	return RuleId(id)
}
func (r RuleId) String() string {
	return uuid.UUID(r).String()
}
