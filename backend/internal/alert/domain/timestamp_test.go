package domain_test

import (
	"testing"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	"github.com/stretchr/testify/assert"
)

func TestTimestamp_NewTimestamp_Valid(t *testing.T) {

	timestamp, err := domain.NewTimestamp("2026-04-02T10:00:00Z")
	assert.NoError(t, err)

	assert.Equal(t, "2026-04-02T10:00:00Z", timestamp.Format(time.RFC3339))
}

func TestTimestamp_NewTimestamp_Invalid(t *testing.T) {

	_, err := domain.NewTimestamp("")
	assert.ErrorIs(t, err, domain.ErrorTimestampInvalid)

}
