package psql

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/psql"
)

type AlertRepository struct {
	queries *psql.Queries
}

func NewAlertRepository(queries *psql.Queries) *AlertRepository {
	return &AlertRepository{queries: queries}
}

func (r *AlertRepository) Save(alert domain.Alert) error {
	return r.queries.InsertAlert(context.Background(), psql.InsertAlertParams{
		Fingerprint: string(alert.Fingerprint),
		Status:      string(alert.Status),
		StartsAt:    alert.StartAt.Time,
		ResolvedAt:  alert.ResolvedAt.Time,
		Labels:      alert.Labels,
		Annotations: alert.Annotations,
	})
}
