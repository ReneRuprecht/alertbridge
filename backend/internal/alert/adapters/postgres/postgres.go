package postgres

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

type AlertRepository struct {
	queries *postgres_db.Queries
}

func NewAlertRepository(queries *postgres_db.Queries) *AlertRepository {
	return &AlertRepository{queries: queries}
}

func (r *AlertRepository) Save(ctx context.Context, alert domain.Alert) error {

	alertRepositoryEntity := toAlertRepositoryEntity(alert)

	return r.queries.InsertAlert(ctx, postgres_db.InsertAlertParams{
		Fingerprint: alertRepositoryEntity.Fingerprint,
		Instance:    alertRepositoryEntity.Instance,
		Status:      alertRepositoryEntity.Status,
		StartsAt:    alertRepositoryEntity.StartsAt,
		ReceivedAt:  alertRepositoryEntity.ReceivedAt,
		Labels:      alertRepositoryEntity.Labels,
		Annotations: alertRepositoryEntity.Annotations,
	})
}

func (r *AlertRepository) ListByInstance(ctx context.Context, instance string) ([]domain.Alert, error) {

	rows, err := r.queries.ListAlertsByInstance(ctx, instance)

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
