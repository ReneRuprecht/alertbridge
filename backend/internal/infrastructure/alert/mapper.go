package alert

import (
	"github.com/reneruprecht/alertbridge/backend/internal/application"
	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

func toListActiveAlertDto(cachedAlerts []application.AlertCacheDto) ListActiveAlertsDto {

	alerts := ListActiveAlertsDto{Alerts: []ActiveAlert{}}

	for _, alert := range cachedAlerts {
		activeAlert := ActiveAlert{
			Fingerprint: alert.Fingerprint,
			AlertName:   alert.AlertName,
			Instance:    alert.Instance,
			Job:         alert.Job,
			Status:      alert.Status,
			StartsAt:    alert.StartsAt,
		}

		alerts.Alerts = append(alerts.Alerts, activeAlert)
	}

	return alerts

}

func toFindAlertsByInstanceDto(historyAlerts []domain.Alert, instance string) FindAlertsByInstanceDto {

	alerts := FindAlertsByInstanceDto{Alerts: []AlertHistory{}}

	for _, alert := range historyAlerts {
		severity := alert.Labels["severity"]
		job := alert.Labels["job"]
		description := alert.Annotations["description"]
		alertName := alert.Labels["alertname"]
		historyAlert := AlertHistory{
			Fingerprint: string(alert.Fingerprint),
			Status:      string(alert.Status),
			Job:         job, Severity: severity,
			AlertName:   alertName,
			Description: description,
			StartAt:     alert.StartAt.Time,
			ReceivedAt:  alert.ReceivedAt.Time,
		}

		alerts.Alerts = append(alerts.Alerts, historyAlert)
	}

	alerts.Instance = instance

	return alerts

}
