package postgres

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

type AlertRepository struct {
	queries *postgres_db.Queries
}

func NewAlertRepository(queries *postgres_db.Queries) *AlertRepository {
	return &AlertRepository{queries: queries}
}

func (r *AlertRepository) Save(alert domain.Alert) error {
	return r.queries.InsertAlert(context.Background(), postgres_db.InsertAlertParams{
		Fingerprint: string(alert.Fingerprint),
		Status:      string(alert.Status),
		StartsAt:    alert.StartAt.Time,
		ResolvedAt:  alert.ResolvedAt.Time,
		Labels:      alert.Labels,
		Annotations: alert.Annotations,
	})
}
