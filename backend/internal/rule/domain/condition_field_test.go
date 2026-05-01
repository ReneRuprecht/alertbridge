package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionField_NewValue_Valid(t *testing.T) {

	field, err := domain.NewConditionField("status")
	require.NoError(t, err)

	assert.Equal(t, "status", string(field))
}

func TestConditionField_NewValue_Empty(t *testing.T) {

	_, err := domain.NewConditionField("")
	require.ErrorIs(t, err, domain.ErrorConditionFieldEmpty)
}

func TestConditionField_NewValue_Invalid(t *testing.T) {

	_, err := domain.NewConditionField("config")
	require.ErrorIs(t, err, domain.ErrorConditionFieldInvalid)
}
