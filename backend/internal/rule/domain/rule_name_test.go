package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRuleName_NewName_Valid(t *testing.T) {

	name, err := domain.NewRuleName("test-rule")
	require.NoError(t, err)

	assert.Equal(t, "test-rule", name.String())
}

func TestRuleName_NewName_Empty(t *testing.T) {

	_, err := domain.NewRuleName("")
	require.ErrorIs(t, err, domain.ErrorRuleNameEmpty)
}
