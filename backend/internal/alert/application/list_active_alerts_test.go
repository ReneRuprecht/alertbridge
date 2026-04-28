package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/stretchr/testify/assert"
)

type mockCacheReader struct {
	err    error
	alerts []application.AlertCacheDto
}

func (m *mockCacheReader) List(context context.Context) ([]application.AlertCacheDto, error) {
	return m.alerts, m.err
}

func TestListAlertsUseCase_Valid(t *testing.T) {

	repo := &mockCacheReader{
		alerts: []application.AlertCacheDto{application.AlertCacheDto{Fingerprint: "fp1", Instance: "backup01", Job: "exporter", Status: "firing", Severity: "ciritical", AlertName: "exporter_down"}},
		err:    nil,
	}
	uc := application.NewListActiveAlertsUseCase(repo)

	ctx := context.Background()
	alerts, err := uc.Execute(ctx)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(alerts))
}

func TestListAlertsUseCase_Error(t *testing.T) {

	repo := &mockCacheReader{
		alerts: []application.AlertCacheDto{},
		err:    errors.New("usecase_error"),
	}
	uc := application.NewListActiveAlertsUseCase(repo)

	ctx := context.Background()
	_, err := uc.Execute(ctx)
	assert.ErrorContains(t, err, "usecase_error")

}
