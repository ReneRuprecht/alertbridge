package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionName_Valid(t *testing.T) {
	data := "send webhook"

	actionName, err := domain.NewActionName(data)

	require.NoError(t, err)
	assert.Equal(t, domain.ActionName("send webhook"), actionName)

}

func TestActionName_EmptyName_ErrorEmtpyName(t *testing.T) {
	data := ""

	_, err := domain.NewActionName(data)

	require.ErrorIs(t, err, domain.ErrorActionNameEmpty)

}

func TestActionName_Spaces_ErrorEmtpyName(t *testing.T) {
	data := "  "

	_, err := domain.NewActionName(data)

	require.ErrorIs(t, err, domain.ErrorActionNameEmpty)

}
