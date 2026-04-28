package alert_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	alert "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_active_alerts"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const url = "/api/v1/alerts"

type mockUseCase struct {
	alerts []application.AlertCacheDto
	err    error
}

func (m *mockUseCase) Execute(ctx context.Context) ([]application.AlertCacheDto, error) {
	return m.alerts, m.err

}

func newListActiveAlertsHandler(alerts []application.AlertCacheDto, error error) http.HandlerFunc {
	uc := &mockUseCase{
		err:    error,
		alerts: alerts,
	}
	handler := alert.HandleListActiveAlerts(uc)
	return handler
}

func newAlertCacheDtos() []application.AlertCacheDto {

	fp1 := "a123"
	status1 := "firing"
	instance1 := "db01"
	job1 := "db-exporter"
	alertName1 := "db down"
	severity1 := "critical"
	startsAt1, _ := time.Parse(time.RFC3339, "2026-01-01T11:00:00Z")

	alert1 := application.AlertCacheDto{Fingerprint: fp1, Instance: instance1, Job: job1, StartsAt: startsAt1, AlertName: alertName1, Status: status1, Severity: severity1}

	fp2 := "a321"
	status2 := "firing"
	instance2 := "web01"
	job2 := "web-exporter"
	alertName2 := "web down"
	severity2 := "critical"
	startsAt2, _ := time.Parse(time.RFC3339, "2026-01-01T12:00:00Z")
	alert2 := application.AlertCacheDto{Fingerprint: fp2, Instance: instance2, Job: job2, StartsAt: startsAt2, AlertName: alertName2, Status: status2, Severity: severity2}

	alerts := []application.AlertCacheDto{alert1, alert2}
	return alerts

}

func TestListAlertsByInstance_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()

	dto := newAlertCacheDtos()
	handler := newListActiveAlertsHandler(dto, nil)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

    var response map[string]any

    err:= json.Unmarshal(w.Body.Bytes(),&response)

    require.NoError(t,err)

    alerts,ok := response["alerts"].([]any)
    require.True(t,ok)
    assert.Len(t,alerts,2)

    firstAlert,ok := alerts[0].(map[string]any)
    require.True(t,ok)
    assert.Equal(t,"a123",firstAlert["fingerprint"])
    assert.Equal(t,"firing",firstAlert["status"])
    assert.Equal(t,"db01",firstAlert["instance"])
    assert.Equal(t,"db-exporter",firstAlert["job"])
    assert.Equal(t,"critical",firstAlert["severity"])
    assert.Equal(t,"2026-01-01T11:00:00Z",firstAlert["startsAt"])

    secondAlert,ok := alerts[1].(map[string]any)
    require.True(t,ok)
    assert.Equal(t,"a321",secondAlert["fingerprint"])
    assert.Equal(t,"firing",secondAlert["status"])
    assert.Equal(t,"web01",secondAlert["instance"])
    assert.Equal(t,"web-exporter",secondAlert["job"])
    assert.Equal(t,"critical",secondAlert["severity"])
    assert.Equal(t,"2026-01-01T12:00:00Z",secondAlert["startsAt"])

}

func TestListAlertsByInstance_InternalServerError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()

	handler := newListActiveAlertsHandler([]application.AlertCacheDto{}, errors.New("failed"))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
