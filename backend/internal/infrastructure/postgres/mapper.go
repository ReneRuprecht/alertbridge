package postgres

import (
	"log"
	"strings"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
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

func toDomain(row postgres_db.FindAlertByInstanceRow) (domain.Alert, error) {

	fp, err := domain.NewFingerprint(row.Fingerprint)

	if err != nil {
		log.Println(err)
		return domain.Alert{}, err
	}
	status, err := domain.NewStatus(row.Status)
	if err != nil {
		log.Println(err)
		return domain.Alert{}, err
	}
	startsAt, err := domain.NewTimestamp(row.StartsAt.Format(time.RFC3339))
	if err != nil {
		log.Println(row.StartsAt.String())
		return domain.Alert{}, err
	}
	resolvedAt, err := domain.NewTimestamp(row.ResolvedAt.Format(time.RFC3339))
	if err != nil {
		log.Println(err)
		return domain.Alert{}, err
	}
	labels := row.Labels
	annotations := row.Annotations

	return domain.Alert{
		Fingerprint: fp,
		Status:      status,
		StartAt:     startsAt,
		ResolvedAt:  resolvedAt,
		Labels:      labels,
		Annotations: annotations,
	}, nil

}
