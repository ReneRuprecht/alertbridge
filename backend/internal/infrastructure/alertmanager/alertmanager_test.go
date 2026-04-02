package alertmanager

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const webhook_url = "/api/v1/alerts"

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

	HandleWebhook(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestHandleWebhook_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, webhook_url, strings.NewReader("{invalid"))
	w := httptest.NewRecorder()

	HandleWebhook(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleWebhook_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, webhook_url, nil)
	w := httptest.NewRecorder()

	HandleWebhook(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", w.Code)
	}
}
