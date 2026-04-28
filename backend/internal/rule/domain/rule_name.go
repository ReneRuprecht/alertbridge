package domain

import "errors"

type RuleName string

var ErrorRuleNameEmpty = errors.New("Name cannot be empty")

func NewRuleName(name string) (RuleName, error) {
	if name == "" {
		return RuleName(""), ErrorRuleNameEmpty
	}

	return RuleName(name), nil
}

func (r RuleName) String() string {
	return string(r)
}
