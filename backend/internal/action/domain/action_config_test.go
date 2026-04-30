package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewActionConfig_Valid(t *testing.T) {

	data := map[string]string{"url": "123"}

	config, err := domain.NewActionConfig(data)

	assert.NoError(t, err)

	assert.Equal(t, "123", config["url"])

}

func TestNewActionConfig_Empty(t *testing.T) {

	data := map[string]string{}

	_, err := domain.NewActionConfig(data)

	assert.ErrorIs(t, err,domain.ErrorActionConfigEmpty)

}
