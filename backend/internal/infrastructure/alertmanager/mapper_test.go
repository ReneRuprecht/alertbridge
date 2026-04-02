package alertmanager

import (
	"errors"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

func TestToDomain_OK(t *testing.T) {

	req := WebhookRequest{
		Alerts: []AlertDto{
			{
				Fingerprint: "123",
				Status:      "firing",
				StartsAt:    "2026-01-01T10:00:00Z",
			},
			{
				Fingerprint: "456",
				Status:      "resolved",
				StartsAt:    "2026-01-01T10:00:00Z",
			},
		},
	}

	alerts, err := toDomain(req)

	if err != nil {
		t.Fatalf("Mapping failed, %v", err)
	}

	if len(alerts) != 2 {
		t.Fatalf("Mapping missed alerts, expected 2 got %d", len(alerts))

	}

}

func TestToDomain_FingerprintEmpty(t *testing.T) {

	req := WebhookRequest{
		Alerts: []AlertDto{
			{
				Fingerprint: "",
				Status:      "firing",
				StartsAt:    "2026-01-01T10:00:00Z",
			},
		},
	}

	_, err := toDomain(req)

	if !errors.Is(err, domain.ErrorFingerprintEmpty) {
		t.Fatalf("expected ErrorEmtpyFingerprint, got %v", err)
	}

}

func TestToDomain_StatusEmpty(t *testing.T) {

	req := WebhookRequest{
		Alerts: []AlertDto{
			{
				Fingerprint: "123",
				Status:      "",
				StartsAt:    "2026-01-01T10:00:00Z",
			},
		},
	}

	_, err := toDomain(req)

	if !errors.Is(err, domain.ErrorStatusEmpty) {
		t.Fatalf("expected ErrorStatusEmpty, got %v", err)
	}

}

func TestToDomain_StatusInvalid(t *testing.T) {

	req := WebhookRequest{
		Alerts: []AlertDto{
			{
				Fingerprint: "123",
				Status:      "active",
				StartsAt:    "2026-01-01T10:00:00Z",
			},
		},
	}

	_, err := toDomain(req)

	if !errors.Is(err, domain.ErrorStatusInvalid) {
		t.Fatalf("expected ErrorStatusInvalid, got %v", err)
	}

}

func TestToDomain_TimestampInvalid(t *testing.T) {

	req := WebhookRequest{
		Alerts: []AlertDto{
			{
				Fingerprint: "123",
				Status:      "firing",
				StartsAt:    "invalid",
			},
		},
	}

	_, err := toDomain(req)

	if !errors.Is(err, domain.ErrorTimestampInvalid) {
		t.Fatalf("expected ErrorTimestampInvalid, got %v", err)
	}

}
