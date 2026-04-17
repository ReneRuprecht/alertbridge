package redis

import (
	"strings"

	"github.com/reneruprecht/alertbridge/backend/internal/application"
	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

func toDto(alert domain.Alert) alertDto {

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

	return alertDto{
		Fingerprint: string(alert.Fingerprint),
		Instance:    instance,
		Job:         job,
		StartsAt:    alert.StartAt.Time,
		AlertName:   alertName,
		Status:      string(alert.Status),
	}
}

func toCacheDto(alert alertDto) application.AlertCacheDto {

	return application.AlertCacheDto{
		Fingerprint: alert.Fingerprint,
		Instance:    alert.Instance,
		Job:         alert.Job,
		Status:      alert.Status,
		StartsAt:    alert.StartsAt,
		AlertName:   alert.AlertName,
	}
}
