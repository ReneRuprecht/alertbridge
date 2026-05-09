package domain

import "errors"

type ConditionField string

var ErrorConditionFieldEmpty = errors.New("Field cannot be empty")
var ErrorConditionFieldInvalid = errors.New("Field is invalid")

const (
	ConditionFieldStatus   ConditionField = "status"
	ConditionFieldSeverity ConditionField = "severity"
)

var validConditionFields = map[ConditionField]struct{}{
	ConditionFieldStatus:   {},
	ConditionFieldSeverity: {},
}

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
	_, ok := validConditionFields[f]
	return ok
}
