package domain

import "strings"

type ActionDescription string

func NewActionDescription(description string) ActionDescription {
	desc := strings.TrimSpace(description)

	return ActionDescription(desc)
}
