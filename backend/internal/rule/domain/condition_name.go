package domain

import "errors"

type ConditionName string

var ErrorConditionNameEmpty = errors.New("Name cannot be empty")

func NewConditionName(name string) (ConditionName, error) {
	if name == "" {
		return ConditionName(""), ErrorConditionNameEmpty
	}

	return ConditionName(name), nil
}
