package domain

import "errors"

type ConditionField string

var ErrorConditionFieldEmpty = errors.New("Field cannot be empty")
var ErrorConditionFieldInvalid = errors.New("Field is invalid")

const ConditionFieldStatus ConditionField = "status"

func NewConditionField(field string) (ConditionField, error) {

	if field == "" {
		return ConditionField(""), ErrorConditionFieldEmpty
	}
	conditionField := ConditionField(field)

	if !conditionField.IsValid() {
		return ConditionField(""), ErrorConditionFieldInvalid
	}

	return conditionField, nil
}

func (f ConditionField) IsValid() bool {
	switch f {
	case ConditionFieldStatus:
		return true
	default:
		return false
	}
}
