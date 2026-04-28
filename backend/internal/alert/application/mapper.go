package application

import (
	"fmt"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

func extractCacheKeyFromAlert(alert domain.Alert) string {

	return fmt.Sprintf("alert:%s", alert.Fingerprint)

}
