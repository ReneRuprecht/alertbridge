package alert_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/application"
	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alert"
)

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

	if instance == "error" {
		return nil, errors.New("Error")
	}

	return []domain.Alert{}, nil
}

func newFindAlertsByInstanceHandler() http.HandlerFunc {
	mockRepo := &MockRepo{}
	uc := application.NewFindAlertsByInstanceUseCase(mockRepo)
	handler := alert.HandleFindAlertsByInstance(uc)
	return handler
}

func newFindAlertsByInstanceUrl(instance string) string {
	var sb strings.Builder
	sb.WriteString("/api/v1/alerts")

	if instance != "" {
		sb.WriteString("?instance=")
		sb.WriteString(instance)
	}

	return sb.String()
}

func TestFindAlertsByInstance_OK(t *testing.T) {
    url := newFindAlertsByInstanceUrl("testing")
	req := httptest.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()

	handler := newFindAlertsByInstanceHandler()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestFindAlertsByInstance_BadRequest(t *testing.T) {
    url := newFindAlertsByInstanceUrl("")

	req := httptest.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()

	handler := newFindAlertsByInstanceHandler()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestFindAlertsByInstance_InternalServerError(t *testing.T) {
    url := newFindAlertsByInstanceUrl("error")

	req := httptest.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()

	handler := newFindAlertsByInstanceHandler()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
