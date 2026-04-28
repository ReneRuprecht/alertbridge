package postgres

import (
	"time"
)

type alertRepositoryEntity struct {
	Fingerprint string
	Instance    string
	Status      string
	ReceivedAt  time.Time
	StartsAt    time.Time
	Labels      map[string]string
	Annotations map[string]string
}
