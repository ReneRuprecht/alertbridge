package application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	"github.com/stretchr/testify/assert"
)

type mockAlertRepoReader struct {
	err    error
	alerts []domain.Alert
}

func (m *mockAlertRepoReader) ListByInstance(context context.Context, instance string) ([]domain.Alert, error) {
	return m.alerts, m.err
}

func TestListAlertsByInstance_Valid(t *testing.T) {
	fp, _ := domain.NewFingerprint("fp1")
	status, _ := domain.NewStatus("firing")
	startAt, _ := domain.NewTimestamp(time.Now().Format(time.RFC3339))
	receivedAt, _ := domain.NewTimestamp(time.Now().Format(time.RFC3339))

	repo := &mockAlertRepoReader{
		alerts: []domain.Alert{domain.Alert{Fingerprint: fp, Status: status, StartAt: startAt, ReceivedAt: receivedAt}},
		err:    nil,
	}
	uc := application.NewListAlertsByInstanceUseCase(repo)

	ctx := context.Background()
	alerts, err := uc.Execute(ctx, "test")

	assert.NoError(t, err)

	assert.Equal(t, 1, len(alerts))
	assert.Equal(t, "fp1", string(alerts[0].Fingerprint))
	assert.Equal(t, "firing", string(alerts[0].Status))

}

func TestListAlertsByInstance_Error(t *testing.T) {
	repo := &mockAlertRepoReader{
		alerts: []domain.Alert{},
		err:    errors.New("db down"),
	}
	uc := application.NewListAlertsByInstanceUseCase(repo)

	ctx := context.Background()
	_, err := uc.Execute(ctx, "test")

	assert.ErrorContains(t, err, "db down")

}
