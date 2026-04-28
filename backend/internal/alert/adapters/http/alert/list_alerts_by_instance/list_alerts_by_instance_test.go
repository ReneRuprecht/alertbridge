package alert_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_alerts_by_instance"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const url = "/api/v1/alerts/"

type mockUseCase struct {
	alerts []domain.Alert
	err    error
}

func (m *mockUseCase) Execute(ctx context.Context, instance string) ([]domain.Alert, error) {
	return m.alerts, m.err

}

func newListAlertsByInstanceHandler(alerts []domain.Alert, error error) http.HandlerFunc {
	uc := &mockUseCase{
		err:    error,
		alerts: alerts,
	}
	handler := alert.HandleListAlertsByInstance(uc)
	return handler
}

func newDomainAlerts() []domain.Alert {

	fp1, _ := domain.NewFingerprint("a123")
	status1, _ := domain.NewStatus("firing")
	startsAt1, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")

	alert1 := domain.Alert{Fingerprint: fp1, Status: status1, StartAt: startsAt1}

	fp2, _ := domain.NewFingerprint("a321")
	status2, _ := domain.NewStatus("resolved")
	startsAt2, _ := domain.NewTimestamp("2026-01-01T11:00:00Z")

	alert2 := domain.Alert{Fingerprint: fp2, Status: status2, StartAt: startsAt2}

	alerts := []domain.Alert{alert1, alert2}

	return alerts
}

func TestListAlertsByInstance_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("instance", "testinstance")

	w := httptest.NewRecorder()

	data := newDomainAlerts()
	handler := newListAlertsByInstanceHandler(data, nil)

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

    var response map[string]any

    err:= json.Unmarshal(w.Body.Bytes(),&response)

    require.NoError(t,err)

    instance := response["instance"]
    assert.Equal(t,"testinstance",instance)

    alerts,ok := response["alerts"].([]any)
    require.True(t,ok)
    assert.Len(t,alerts,2)

    firstAlert,ok := alerts[0].(map[string]any)
    require.True(t,ok)
    assert.Equal(t,"a123",firstAlert["fingerprint"])
    assert.Equal(t,"firing",firstAlert["status"])
    assert.Equal(t,"2026-01-01T10:00:00Z",firstAlert["startsAt"])

    secondAlert,ok := alerts[1].(map[string]any)
    require.True(t,ok)
    assert.Equal(t,"a321",secondAlert["fingerprint"])
    assert.Equal(t,"resolved",secondAlert["status"])
    assert.Equal(t,"2026-01-01T11:00:00Z",secondAlert["startsAt"])
}

func TestListAlertsByInstance_BadRequest(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()

	alerts := []domain.Alert{}
	handler := newListAlertsByInstanceHandler(alerts, nil)

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestListAlertsByInstance_InternalServerError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("instance", "testinstance")

	w := httptest.NewRecorder()

	err := errors.New("failed")
	handler := newListAlertsByInstanceHandler(nil, err)

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
