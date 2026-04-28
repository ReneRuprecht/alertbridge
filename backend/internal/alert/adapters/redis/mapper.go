package redis

import (
	"strings"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

func toAlertCacheEntity(alert domain.Alert) alertCacheEntity {

	instance := "unknown"

	inst := strings.Split(alert.Labels["instance"], ":")[0]
	if inst != "" {
		instance = inst
	}

	job := alert.Labels["job"]
	if job == "" {
		job = "unknown"
	}

	alertName := "unknown"

	alertNameLabel := alert.Labels["alertname"]
	if alertNameLabel != "" {
		alertName = alertNameLabel
	}

	severity := alert.Labels["severity"]
	if severity == "" {
		severity = "unknown"
	}

	return alertCacheEntity{
		Fingerprint: string(alert.Fingerprint),
		Instance:    instance,
		Job:         job,
		StartsAt:    alert.StartAt.Time,
		AlertName:   alertName,
		Status:      string(alert.Status),
		Severity:    severity,
	}
}

func toCacheDto(entity alertCacheEntity) application.AlertCacheDto {

	return application.AlertCacheDto{
		Fingerprint: entity.Fingerprint,
		Instance:    entity.Instance,
		Job:         entity.Job,
		Status:      entity.Status,
		StartsAt:    entity.StartsAt,
		AlertName:   entity.AlertName,
		Severity:    entity.Severity,
	}
}
