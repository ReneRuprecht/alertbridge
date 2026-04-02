package domain

type Alert struct {
	Fingerprint Fingerprint
	Status      Status
	Labels      map[string]string
	Annotations map[string]string
	StartAt     Timestamp
}
