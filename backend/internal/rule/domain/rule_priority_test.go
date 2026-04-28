package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRulePriority_NewPriority_Valid(t *testing.T) {

	priority, err := domain.NewRulePriority(1)
	require.NoError(t, err)

	assert.Equal(t, 1, priority.Int())
}

func TestRulePriority_NewPriority_ErrorNegative(t *testing.T) {

	_, err := domain.NewRulePriority(-1)
	require.ErrorIs(t, err, domain.ErrorRulePriorityNegative)
}
