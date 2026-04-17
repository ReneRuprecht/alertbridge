package alert

import (
	"time"
)

type ListActiveAlertsDto struct {
	Alerts []ActiveAlert `json:"alerts"`
}

type ActiveAlert struct {
	Fingerprint string    `json:"fingerprint"`
	Instance    string    `json:"instance"`
	Status      string    `json:"status"`
	StartsAt    time.Time `json:"startsAt"`
	AlertName   string    `json:"alertName"`
}
