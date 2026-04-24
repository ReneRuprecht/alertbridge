package postgres

import (
	"time"
)

type AlertDto struct {
	Fingerprint string
	Instance    string
	Status      string
	ReceivedAt  time.Time
	StartsAt    time.Time
	Labels      map[string]string
	Annotations map[string]string
}
