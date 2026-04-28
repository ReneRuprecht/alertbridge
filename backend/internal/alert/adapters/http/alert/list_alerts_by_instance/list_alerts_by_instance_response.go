package alert

import "time"

type AlertHistory struct {
	Fingerprint string    `json:"fingerprint"`
	Status      string    `json:"status"`
	AlertName   string    `json:"alertName"`
	Job         string    `json:"job"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"`
	StartAt     time.Time `json:"startsAt"`
	ReceivedAt  time.Time `json:"receivedAt"`
}

type listAlertsByInstanceResponse struct {
	Instance string         `json:"instance"`
	Alerts   []AlertHistory `json:"alerts"`
}
