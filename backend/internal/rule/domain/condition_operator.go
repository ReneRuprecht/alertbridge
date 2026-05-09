package domain

import "errors"

type ConditionOperator string

var ErrorConditionOperatorEmpty = errors.New("Operator cannot be empty")
var ErrorConditionOperatorInvalid = errors.New("Operator is invalid")

const ConditionOperatorEquals ConditionOperator = "equals"

func NewConditionOperator(operator string) (ConditionOperator, error) {

	if operator == "" {
		return ConditionOperator(""), ErrorConditionOperatorEmpty
	}
	conditionOperator := ConditionOperator(operator)

	if !conditionOperator.IsValid() {
		return ConditionOperator(""), ErrorConditionOperatorInvalid
	}

	return conditionOperator, nil
}

func (f ConditionOperator) IsValid() bool {
	switch f {
	case ConditionOperatorEquals:
		return true
	default:
		return false
	}
}
func (o ConditionOperator) Apply(a, b string) bool {
	switch o {
	case ConditionOperatorEquals:
		return a == b
	}
	return false
}
