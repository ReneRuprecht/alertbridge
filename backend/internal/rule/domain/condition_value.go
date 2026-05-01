package domain

import "errors"

type ConditionValue string

var ErrorConditionValueEmpty = errors.New("Value cannot be empty")

func NewConditionValue(value string) (ConditionValue, error) {
	if value == "" {
		return ConditionValue(""), ErrorConditionValueEmpty
	}

	return ConditionValue(value), nil
}
