package findalertbyfingerprint

import (
	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

func toAlertFingerprint(fingerprint string) (domain.Fingerprint, error) {
	return domain.NewFingerprint(fingerprint)
}

func toFingerAlertByFingerprintResponse(alert application.AlertCacheDto) FindAlertByFingerprintResponse {

	return FindAlertByFingerprintResponse{
		Fingerprint: alert.Fingerprint,
		AlertName:   alert.AlertName,
		Instance:    alert.Instance,
		Job:         alert.Job,
		Status:      alert.Status,
		Severity:    alert.Severity,
		StartsAt:    alert.StartsAt,
	}
}
