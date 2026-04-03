package redis

import "time"

type alertDto struct {
	Fingerprint string    `json:"fingerprint"`
	Instance    string    `json:"instance"`
	StartsAt    time.Time `json:"starts_at"`
	AlertName   string    `json:"alert_name"`
	Status      string    `json:"status"`
}
