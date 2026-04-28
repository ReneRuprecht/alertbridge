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

type mockAlertRepositoryWriter struct {
	err        error
	saveCalled bool
}

func (m *mockAlertRepositoryWriter) Save(context context.Context, alert domain.Alert) error {
	m.saveCalled = true
	return m.err
}

type mockAlertCacheWriter struct {
	err          error
	saveCalled   bool
	deleteCalled bool
}

func (m *mockAlertCacheWriter) Save(context context.Context, key string, alert domain.Alert) error {
	m.saveCalled = true
	return m.err
}

func (m *mockAlertCacheWriter) DeleteByKey(context context.Context, key string) error {
	m.deleteCalled = true
	return m.err
}

func setupUsecase(repoError error, cacheError error) (*mockAlertRepositoryWriter, *mockAlertCacheWriter, application.SaveAlertsWithCacheUseCase) {

	repo := &mockAlertRepositoryWriter{
		saveCalled: false,
		err:        repoError,
	}
	cache := &mockAlertCacheWriter{
		saveCalled:   false,
		deleteCalled: false,
		err:          cacheError,
	}
	uc := application.NewSaveAlertsWithCacheUseCase(repo, cache)

	return repo, cache, uc

}

func getFiringDomainAlerts() []domain.Alert {

	fp, _ := domain.NewFingerprint("fp1")
	status, _ := domain.NewStatus("firing")
	startAt, _ := domain.NewTimestamp(time.Now().Format(time.RFC3339))
	receivedAt, _ := domain.NewTimestamp(time.Now().Format(time.RFC3339))
	data := []domain.Alert{domain.Alert{Fingerprint: fp, Status: status, StartAt: startAt, ReceivedAt: receivedAt}}
	return data
}
func getResolvedDomainAlerts() []domain.Alert {

	fp, _ := domain.NewFingerprint("fp1")
	status, _ := domain.NewStatus("resolved")
	startAt, _ := domain.NewTimestamp(time.Now().Format(time.RFC3339))
	receivedAt, _ := domain.NewTimestamp(time.Now().Format(time.RFC3339))
	data := []domain.Alert{domain.Alert{Fingerprint: fp, Status: status, StartAt: startAt, ReceivedAt: receivedAt}}
	return data
}

func TestSaveAlertsWithCache_Valid(t *testing.T) {

	repo, cache, uc := setupUsecase(nil, nil)

	data := getFiringDomainAlerts()
	ctx := context.Background()
	err := uc.Execute(ctx, data)

	assert.NoError(t, err)

	assert.Equal(t, true, repo.saveCalled)
	assert.Equal(t, true, cache.saveCalled)
	assert.Equal(t, false, cache.deleteCalled)

}

func TestSaveAlertsWithCache_RepoError(t *testing.T) {

	repo, cache, uc := setupUsecase(errors.New("repo error"), nil)

	data := getFiringDomainAlerts()

	ctx := context.Background()
	err := uc.Execute(ctx, data)

	assert.ErrorContains(t, err, "repo error")

	assert.Equal(t, true, repo.saveCalled)
	assert.Equal(t, false, cache.saveCalled)
	assert.Equal(t, false, cache.deleteCalled)

}

func TestSaveAlertsWithCache_CacheSaveError(t *testing.T) {

	repo, cache, uc := setupUsecase(nil, errors.New("cache error"))

	data := getFiringDomainAlerts()

	ctx := context.Background()
	err := uc.Execute(ctx, data)

	assert.ErrorContains(t, err, "cache error")

	assert.Equal(t, true, repo.saveCalled)
	assert.Equal(t, true, cache.saveCalled)
	assert.Equal(t, false, cache.deleteCalled)

}

func TestSaveAlertsWithCache_CacheDeleteError(t *testing.T) {

	repo, cache, uc := setupUsecase(nil, errors.New("cache error"))

	data := getResolvedDomainAlerts()

	ctx := context.Background()
	err := uc.Execute(ctx, data)

	assert.ErrorContains(t, err, "cache error")
	assert.Equal(t, true, repo.saveCalled)
	assert.Equal(t, false, cache.saveCalled)
	assert.Equal(t, true, cache.deleteCalled)

}
