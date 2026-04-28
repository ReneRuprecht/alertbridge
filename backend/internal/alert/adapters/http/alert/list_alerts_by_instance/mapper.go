package alert

import (
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

func toListAlertsByInstanceResponse(historyAlerts []domain.Alert, instance string) listAlertsByInstanceResponse {

	alerts := listAlertsByInstanceResponse{Alerts: []AlertHistory{}}

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
