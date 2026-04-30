package rabbitmq

import "github.com/reneruprecht/alertbridge/backend/internal/alert/domain"

func toAlertFingerprint(fingerprint string) (domain.Fingerprint, error) {

	return domain.NewFingerprint(fingerprint)

}
