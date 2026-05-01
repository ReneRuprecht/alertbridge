package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionOperator_NewValue_Valid(t *testing.T) {

	operator, err := domain.NewConditionOperator("equals")
	require.NoError(t, err)

	assert.Equal(t, "equals", string(operator))
}

func TestConditionOperator_NewValue_Empty(t *testing.T) {

	_, err := domain.NewConditionOperator("")
	require.ErrorIs(t, err, domain.ErrorConditionOperatorEmpty)
}

func TestConditionOperator_NewValue_Invalid(t *testing.T) {

	_, err := domain.NewConditionOperator("greater")
	require.ErrorIs(t, err, domain.ErrorConditionOperatorInvalid)
}
