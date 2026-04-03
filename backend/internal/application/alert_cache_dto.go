package application

import "time"

type AlertCacheDto struct {
	Fingerprint string
	Instance    string
	StartsAt    time.Time
	AlertName   string
	Status      string
}
