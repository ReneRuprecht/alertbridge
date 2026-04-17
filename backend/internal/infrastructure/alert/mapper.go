package alert

import (
	"github.com/reneruprecht/alertbridge/backend/internal/application"
)

func toListActiveAlertDto(cachedAlerts []application.AlertCacheDto) ListActiveAlertsDto {

	alerts := ListActiveAlertsDto{Alerts: []ActiveAlert{}}

	for _, alert := range cachedAlerts {
		activeAlert := ActiveAlert{Fingerprint: alert.Fingerprint, AlertName: alert.AlertName, Instance: alert.Instance, Status: alert.Status, StartsAt: alert.StartsAt}

		alerts.Alerts = append(alerts.Alerts, activeAlert)
	}

	return alerts

}
