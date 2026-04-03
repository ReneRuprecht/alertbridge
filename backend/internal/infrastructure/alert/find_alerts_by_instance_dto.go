package alert

import "github.com/reneruprecht/alertbridge/backend/internal/domain"

type FindAlertsByInstanceDto struct {
	Instance string         `json:"instance"`
	Alerts   []domain.Alert `json:"alerts"`
}
