package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionValue_NewValue_Valid(t *testing.T) {

	value, err := domain.NewConditionValue("dev")
	require.NoError(t, err)

	assert.Equal(t, "dev", string(value))
}

func TestConditionValue_NewValue_Empty(t *testing.T) {

	_, err := domain.NewConditionValue("")
	require.ErrorIs(t, err, domain.ErrorConditionValueEmpty)
}
