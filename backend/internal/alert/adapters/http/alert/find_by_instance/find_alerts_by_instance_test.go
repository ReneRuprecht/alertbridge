package alert_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/find_by_instance"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

const url = "/api/v1/alerts/"

type mockUseCase struct {
	alerts []domain.Alert
	err    error
}

func (m *mockUseCase) Execute(ctx context.Context, instance string) ([]domain.Alert, error) {
	return m.alerts, m.err

}

func newFindAlertsByInstanceHandler(alerts []domain.Alert, error error) http.HandlerFunc {
	uc := &mockUseCase{
		err:    error,
		alerts: alerts,
	}
	handler := alert.HandleFindAlertsByInstance(uc)
	return handler
}

func TestFindAlertsByInstance_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("instance", "testinstance")

	w := httptest.NewRecorder()

	uc := &mockUseCase{
		alerts: []domain.Alert{},
	}
	handler := alert.HandleFindAlertsByInstance(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestFindAlertsByInstance_BadRequest(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()

	uc := &mockUseCase{
		alerts: []domain.Alert{},
	}
	handler := alert.HandleFindAlertsByInstance(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestFindAlertsByInstance_InternalServerError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("instance", "testinstance")

	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: errors.New("failed"),
	}
	handler := alert.HandleFindAlertsByInstance(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
