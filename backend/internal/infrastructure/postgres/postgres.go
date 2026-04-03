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

func (r *AlertRepository) Save(ctx context.Context, alert domain.Alert) error {

	alertDto := toDto(alert)

	return r.queries.InsertAlert(ctx, postgres_db.InsertAlertParams{
		Fingerprint: alertDto.Fingerprint,
		Instance:    alertDto.Instance,
		Status:      alertDto.Status,
		StartsAt:    alertDto.StartsAt,
		ResolvedAt:  alertDto.ResolvedAt,
		Labels:      alertDto.Labels,
		Annotations: alertDto.Annotations,
	})
}

func (r *AlertRepository) FindAlertsByInstance(ctx context.Context, instance string) ([]domain.Alert, error) {

	rows, err := r.queries.FindAlertByInstance(ctx, instance)

	if err != nil {
		return nil, err
	}

	alerts := make([]domain.Alert, len(rows))

	for i, row := range rows {
		alert, err := toDomain(row)

		if err != nil {
			return nil, err
		}

		alerts[i] = alert

	}

	return alerts, nil
}
