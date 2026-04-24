package domain

import "errors"

type RulePriority int

var ErrorRulePriorityNegative = errors.New("Rule priority cannot be negative")

func NewRulePriority(priority int) (RulePriority, error) {
	if priority < 0 {
		return RulePriority(-1), ErrorRulePriorityNegative
	}

	return RulePriority(priority), nil
}

func (p RulePriority) Int() int {
	return int(p)
}
