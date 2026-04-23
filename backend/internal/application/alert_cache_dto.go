package application

import "time"

type AlertCacheDto struct {
	Fingerprint string
	Instance    string
	Job         string
	StartsAt    time.Time
	AlertName   string
	Status      string
	Severity    string
}
