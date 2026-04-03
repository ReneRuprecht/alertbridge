package postgres

import (
	"strings"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

func toDto(alert domain.Alert) AlertDto {

	instance := "unknown"

	inst := strings.Split(alert.Labels["instance"], ":")[0]
	if inst != "" {
		instance = inst
	}

	return AlertDto{
		Fingerprint: string(alert.Fingerprint),
		Instance:    instance,
		Status:      string(alert.Status),
		StartsAt:    alert.StartAt.Time,
		ResolvedAt:  alert.ResolvedAt.Time,
		Labels:      alert.Labels,
		Annotations: alert.Annotations,
	}
}
