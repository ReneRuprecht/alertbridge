package rule

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

const url = "/api/v1/rule"

type mockUseCase struct {
	err error
}

func (m *mockUseCase) Execute(ctx context.Context, rule domain.Rule) error {
	return m.err
}

func TestHandleCreateRule_Created(t *testing.T) {
	body := `{
        "name": "backend-critical-alerts",
        "description": "Handles critical alerts for backend services",
        "priority": 100,
        "enabled": true
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := HandleCreateRule(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}

func TestHandleCreateRule_InvalidJson(t *testing.T) {
	body := `{
        "name": "backend-critical-alerts"
        "description": "Handles critical alerts for backend services",
        "priority": 100,
        "enabled": true
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := HandleCreateRule(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleCreateRule_InvalidRequest(t *testing.T) {
	body := `{
        "name": "",
        "description": "Handles critical alerts for backend services",
        "priority": 100,
        "enabled": true
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := HandleCreateRule(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleCreateRule_InternalServerError(t *testing.T) {
	body := `{
        "name": "backend-critical-alerts",
        "description": "Handles critical alerts for backend services",
        "priority": 100,
        "enabled": true
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: errors.New("failed"),
	}
	handler := HandleCreateRule(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
