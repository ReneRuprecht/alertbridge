package domain

import (
	"errors"
	"strings"
)

type ActionName string

var ErrorActionNameEmpty = errors.New("name cannot be empty")

func NewActionName(name string) (ActionName, error) {
	actionName := strings.TrimSpace(name)
	if actionName == "" {
		return "", ErrorActionNameEmpty
	}

	return ActionName(actionName), nil
}
