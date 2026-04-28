package redis

import (
	"testing"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	"github.com/stretchr/testify/assert"
)

func TestToDto_UnknownInstance(t *testing.T) {

	fp, _ := domain.NewFingerprint("z123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")
	labels := make(map[string]string)
	labels["job"] = "node_exporter"
	labels["severity"] = "critical"

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt, Labels: labels}

	entity := toAlertCacheEntity(alert)

	assert.Equal(t, "unknown", entity.Instance)
	assert.Equal(t, "z123", entity.Fingerprint)
	assert.Equal(t, "firing", entity.Status)
	assert.Equal(t, "critical", entity.Severity)
	assert.Equal(t, "node_exporter", entity.Job)
	assert.Equal(t, "2026-01-01T10:00:00Z", entity.StartsAt.Format(time.RFC3339))
}

func TestToDto_KnownInstance(t *testing.T) {

	fp, _ := domain.NewFingerprint("z123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")
	labels := make(map[string]string)
	labels["instance"] = "testinstance"

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt, Labels: labels}

	entity := toAlertCacheEntity(alert)

	assert.Equal(t, "testinstance", entity.Instance)
}

func TestToCacheDto_Valid(t *testing.T) {
	start, _ := time.Parse(time.RFC3339, "2026-01-01T10:00:00Z")

	alertCacheEntity := alertCacheEntity{
		Fingerprint: "z123",
		Instance:    "testinstance",
		AlertName:   "InstanceDown",
		Status:      "firing",
		StartsAt:    start,
		Severity:    "critical",
	}

	alertCacheDto := toCacheDto(alertCacheEntity)

	assert.Equal(t, alertCacheEntity.Fingerprint, alertCacheDto.Fingerprint)
	assert.Equal(t, alertCacheEntity.Instance, alertCacheDto.Instance)
	assert.Equal(t, alertCacheEntity.StartsAt, alertCacheDto.StartsAt)
	assert.Equal(t, alertCacheEntity.AlertName, alertCacheDto.AlertName)
	assert.Equal(t, alertCacheEntity.Status, alertCacheDto.Status)
	assert.Equal(t, alertCacheEntity.Severity, alertCacheDto.Severity)

}
