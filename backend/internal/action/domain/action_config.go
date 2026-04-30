package domain

import (
	"errors"
	"maps"
)

type ActionConfig map[string]string

var ErrorActionConfigEmpty = errors.New("config cannot be emtpy")

func NewActionConfig(config map[string]string) (ActionConfig, error) {

	if len(config) == 0 {
		return nil, ErrorActionConfigEmpty
	}

	c := make(map[string]string, len(config))

	maps.Copy(c, config)

	return ActionConfig(c), nil
}
