package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewActionType_Valid(t *testing.T) {
	data := "webhook"

	actionType, err := domain.NewActionType(data)
	require.NoError(t, err)
	assert.NotEqual(t, domain.ActionType(""), actionType)
}

func TestNewActionType_Invalid(t *testing.T) {
	data := "bird"

	_, err := domain.NewActionType(data)
	require.Error(t, err)
	assert.ErrorContains(t, err, "invalid action type: bird")
}
