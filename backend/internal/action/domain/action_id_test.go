package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewActionID_Valid(t *testing.T) {
	id, err := domain.NewActionID()

	assert.NoError(t, err)
	assert.NotEqual(t, domain.ActionID{}, id)
}
