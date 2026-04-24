package alertmanager

import (
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

func toDomain(req WebhookRequest) ([]domain.Alert, error) {
	alerts := make([]domain.Alert, len(req.Alerts))

	for i, alert := range req.Alerts {

		fingerprint, err := domain.NewFingerprint(alert.Fingerprint)

		if err != nil {
			return nil, err
		}

		status, err := domain.NewStatus(alert.Status)

		if err != nil {
			return nil, err
		}

		startsAt, err := domain.NewTimestamp(alert.StartsAt)
		if err != nil {
			return nil, err
		}

		receivedAt, err := domain.NewTimestamp(time.Now().Format(time.RFC3339))
		if err != nil {
			return nil, err
		}

		alerts[i] = domain.Alert{
			Fingerprint: fingerprint,
			Status:      status,
			StartAt:     startsAt,
			ReceivedAt:  receivedAt,
			Labels:      alert.Labels,
			Annotations: alert.Annotations,
		}

	}

	return alerts, nil

}
