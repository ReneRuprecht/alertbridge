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

type mockUseCase struct {
	err error
}

func (m *mockUseCase) Execute(ctx context.Context, alerts []domain.Alert) error {
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

	uc := &mockUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestHandleWebhook_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, webhook_url, strings.NewReader("{invalid"))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleWebhook_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, webhook_url, nil)
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", w.Code)
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

	uc := &mockUseCase{
		err: nil,
	}
	handler := alertmanager.HandleWebhook(uc)
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

	uc := &mockUseCase{
		err: errors.New("failed"),
	}
	handler := alertmanager.HandleWebhook(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
