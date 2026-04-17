package redis

import (
	"testing"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestToDto_UnknownInstance(t *testing.T) {

	fp, _ := domain.NewFingerprint("z123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")
	labels := make(map[string]string)
	labels["job"] = "node_exporter"

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt, Labels: labels}

	dto := toDto(alert)

	assert.Equal(t, "unknown", dto.Instance)
	assert.Equal(t, "z123", dto.Fingerprint)
	assert.Equal(t, "firing", dto.Status)
	assert.Equal(t, "node_exporter", dto.Job)
	assert.Equal(t, "2026-01-01T10:00:00Z", dto.StartsAt.Format(time.RFC3339))
}

func TestToDto_KnownInstance(t *testing.T) {

	fp, _ := domain.NewFingerprint("z123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")
	labels := make(map[string]string)
	labels["instance"] = "testinstance"

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt, Labels: labels}

	dto := toDto(alert)

	assert.Equal(t, "testinstance", dto.Instance)
}

func TestToCacheDto_Valid(t *testing.T) {
	start, _ := time.Parse(time.RFC3339, "2026-01-01T10:00:00Z")

	alertDto := alertDto{
		Fingerprint: "z123",
		Instance:    "testinstance",
		AlertName:   "InstanceDown",
		Status:      "firing",
		StartsAt:    start,
	}

	alertCacheDto := toCacheDto(alertDto)

	assert.Equal(t, alertDto.Fingerprint, alertCacheDto.Fingerprint)
	assert.Equal(t, alertDto.Instance, alertCacheDto.Instance)
	assert.Equal(t, alertDto.StartsAt, alertCacheDto.StartsAt)
	assert.Equal(t, alertDto.AlertName, alertCacheDto.AlertName)
	assert.Equal(t, alertDto.Status, alertCacheDto.Status)

}
