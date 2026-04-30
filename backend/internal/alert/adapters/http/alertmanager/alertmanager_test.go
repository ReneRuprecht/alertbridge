package alertmanager_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alertmanager"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

const webhook_url = "/api/v1/alertmanager"

type mockSaveAlertsWithCacheUseCase struct {
	err error
}

func (m *mockSaveAlertsWithCacheUseCase) Execute(ctx context.Context, alerts []domain.Alert) error {
	return m.err
}

type mockPublishAlertsUseCase struct {
	err error
}

func (m *mockPublishAlertsUseCase) Execute(ctx context.Context, alerts []domain.Alert) error {
	return m.err
}

func TestHandleWebhook_OK(t *testing.T) {
	body := `{
        "alerts": [{
            "status": "firing",
            "fingerprint": "abc",
            "startsAt": "2026-04-02T10:00:00Z",
            "annotations": {
                "team": "dev"
            },
            "labels": {
                "env": "dev"
            }
        }]
    }`

	req := httptest.NewRequest(http.MethodPost, webhook_url, strings.NewReader(body))
	w := httptest.NewRecorder()

	saveAlertsWithCacheUseCase := &mockSaveAlertsWithCacheUseCase{
		err: nil,
	}
	publishAlertsUseCase := &mockPublishAlertsUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(saveAlertsWithCacheUseCase, publishAlertsUseCase)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestHandleWebhook_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, webhook_url, strings.NewReader("{invalid"))
	w := httptest.NewRecorder()

	saveAlertsWithCacheUseCase := &mockSaveAlertsWithCacheUseCase{
		err: nil,
	}
	publishAlertsUseCase := &mockPublishAlertsUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(saveAlertsWithCacheUseCase, publishAlertsUseCase)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleWebhook_InvalidRequest(t *testing.T) {
	body := `{
        "alerts": [{
            "status": "",
            "fingerprint": "abc",
            "startsAt": "2026-04-02T10:00:00Z",
            "annotations": {
                "team": "dev"
            },
            "labels": {
                "env": "dev"
            }
        }]
    }`

	req := httptest.NewRequest(http.MethodPost, webhook_url, strings.NewReader(body))
	w := httptest.NewRecorder()

	saveAlertsWithCacheUseCase := &mockSaveAlertsWithCacheUseCase{
		err: nil,
	}
	publishAlertsUseCase := &mockPublishAlertsUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(saveAlertsWithCacheUseCase, publishAlertsUseCase)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 405, got %d", w.Code)
	}
}

func TestHandleWebhook_InternalServerError(t *testing.T) {
	body := `{
        "alerts": [{
            "status": "firing",
            "fingerprint": "abc",
            "startsAt": "2026-04-02T10:00:00Z",
            "annotations": {
                "team": "dev"
            },
            "labels": {
                "env": "dev"
            }
        }]
    }`

	req := httptest.NewRequest(http.MethodPost, webhook_url, strings.NewReader(body))
	w := httptest.NewRecorder()

	saveAlertsWithCacheUseCase := &mockSaveAlertsWithCacheUseCase{
		err: errors.New("failed"),
	}
	publishAlertsUseCase := &mockPublishAlertsUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(saveAlertsWithCacheUseCase, publishAlertsUseCase)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
