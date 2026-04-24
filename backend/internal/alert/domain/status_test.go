package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	"github.com/stretchr/testify/assert"
)

func TestStatus_NewStatus_Valid(t *testing.T) {

	status, err := domain.NewStatus("firing")
	assert.NoError(t, err)

	assert.Equal(t, "firing", string(status))
}

func TestStatus_NewStatus_Empty(t *testing.T) {

	_, err := domain.NewStatus("")
	assert.ErrorIs(t, err, domain.ErrorStatusEmpty)

}

func TestStatus_NewStatus_Invalid(t *testing.T) {

	_, err := domain.NewStatus("done")
	assert.ErrorIs(t, err, domain.ErrorStatusInvalid)

}
