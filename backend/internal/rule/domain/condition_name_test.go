package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionName_NewValue_Valid(t *testing.T) {

	name, err := domain.NewConditionName("equals")
	require.NoError(t, err)

	assert.Equal(t, "equals", string(name))
}

func TestConditionName_NewValue_Empty(t *testing.T) {

	_, err := domain.NewConditionName("")
	require.ErrorIs(t, err, domain.ErrorConditionNameEmpty)
}
