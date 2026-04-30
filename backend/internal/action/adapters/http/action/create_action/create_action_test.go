package createaction

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
)

const url = "/api/v1/actions"

type mockUseCase struct {
	err error
}

func (m *mockUseCase) Execute(ctx context.Context, rule domain.Action) error {
	return m.err
}

func TestHandleCreateAction_Created(t *testing.T) {
	body := `{
        "name": "slack sender",
        "description": "sends slack message",
        "type": "webhook",
        "config": {
          "url": "slack",
          "channel": "alert"
        }
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := HandleCreateAction(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}

func TestHandleCreateAction_MissingName(t *testing.T) {
	body := `{
        "description": "sends slack message",
        "type": "webhook",
        "config": {
          "url": "slack",
          "channel": "alert"
        }
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := HandleCreateAction(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
