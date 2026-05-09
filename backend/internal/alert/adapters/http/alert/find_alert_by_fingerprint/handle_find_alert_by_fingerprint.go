package findalertbyfingerprint

import (
	"encoding/json"
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
)

func HandleFindAlertByFingerprint(uc application.FindAlertByFingerprintUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		fingerprint := r.PathValue("fingerprint")
		if fingerprint == "" {
			http.Error(w, "fingerprint is required", http.StatusBadRequest)
		}

		alertFingerprint, err := toAlertFingerprint(fingerprint)

		ctx := r.Context()
		alert, err := uc.Execute(ctx, alertFingerprint)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := toFingerAlertByFingerprintResponse(alert)
		json.NewEncoder(w).Encode(response)
	}
}
