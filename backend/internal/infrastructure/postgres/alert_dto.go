package postgres

import (
	"time"
)

type AlertDto struct {
	Fingerprint string
	Instance    string
	Status      string
	StartsAt    time.Time
	ResolvedAt  time.Time
	Labels      map[string]string
	Annotations map[string]string
}
