package alert

import (
	"github.com/reneruprecht/alertbridge/backend/internal/application"
)

type ListActiveAlertsDto struct {
	Alerts []application.AlertCacheDto `json:"alerts"`
}
