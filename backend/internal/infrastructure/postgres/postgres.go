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

	alertDto := toDto(alert)

	return r.queries.InsertAlert(context.Background(), postgres_db.InsertAlertParams{
		Fingerprint: alertDto.Fingerprint,
		Instance:    alertDto.Instance,
		Status:      alertDto.Status,
		StartsAt:    alertDto.StartsAt,
		ResolvedAt:  alertDto.ResolvedAt,
		Labels:      alertDto.Labels,
		Annotations: alertDto.Annotations,
	})
}
