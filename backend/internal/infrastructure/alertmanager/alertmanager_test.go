package alertmanager_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/application"
	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alertmanager"
)

const webhook_url = "/api/v1/alertmanager"

type MockRepo struct {
	Saved []domain.Alert
	Err   error
}

func (m *MockRepo) Save(ctx context.Context, a domain.Alert) error {
	if m.Err != nil {
		return m.Err
	}
	m.Saved = append(m.Saved, a)
	return nil
}

func (m *MockRepo) FindAlertsByInstance(ctx context.Context, instance string) ([]domain.Alert, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	return nil, nil
}

func newHandleWebhook() http.HandlerFunc {
	mockRepo := &MockRepo{}
	uc := application.NewReceiveAlertUseCase(mockRepo)
	handler := alertmanager.HandleWebhook(uc)
	return handler
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

	handler := newHandleWebhook()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestHandleWebhook_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, webhook_url, strings.NewReader("{invalid"))
	w := httptest.NewRecorder()

	handler := newHandleWebhook()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleWebhook_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, webhook_url, nil)
	w := httptest.NewRecorder()

	handler := newHandleWebhook()
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

	handler := newHandleWebhook()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 405, got %d", w.Code)
	}
}
