package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewActionDescription_Description_Valid(t *testing.T) {
	data := "sends slack message"

	desc := domain.NewActionDescription(data)

	assert.Equal(t, domain.ActionDescription("sends slack message"), desc)
}

func TestNewActionDescription_EmptyDescription_Valid(t *testing.T) {
	data := ""

	desc := domain.NewActionDescription(data)

	assert.Equal(t, domain.ActionDescription(""), desc)
}
