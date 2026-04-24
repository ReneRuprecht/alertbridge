package domain_test

import (
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	"github.com/stretchr/testify/assert"
)

func TestFingerprint_NewFingerprint_Valid(t *testing.T) {

	fp, err := domain.NewFingerprint("fp1")
	assert.NoError(t, err)

	assert.Equal(t, "fp1", string(fp))
}

func TestFingerprint_NewFingerprint_Empty(t *testing.T) {

	_, err := domain.NewFingerprint("")
	assert.ErrorIs(t, err, domain.ErrorFingerprintEmpty)

}
