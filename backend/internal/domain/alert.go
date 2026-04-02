package domain

import "time"

type Fingerprint string


type Timestamp struct {
	time.Time
}

type Alert struct {
	Fingerprint Fingerprint
	Status      Status
	Labels      map[string]string
	Annotations map[string]string
	StartAt     Timestamp
}
