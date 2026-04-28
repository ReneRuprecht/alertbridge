package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type SaveAlertsWithCacheUseCase interface {
	Execute(ctx context.Context, alerts []domain.Alert) error
}

type saveAlertsWithCacheUseCase struct {
	repo  AlertRepositoryWriter
	cache AlertCacheWriter
}

func NewSaveAlertsWithCacheUseCase(repo AlertRepositoryWriter, cache AlertCacheWriter) *saveAlertsWithCacheUseCase {
	return &saveAlertsWithCacheUseCase{repo: repo, cache: cache}
}

func (uc *saveAlertsWithCacheUseCase) Execute(ctx context.Context, alerts []domain.Alert) error {
	for _, alert := range alerts {

		if err := uc.handleSaveAlertsWithCache(ctx, alert); err != nil {
			return err
		}

	}
	return nil
}

func (uc *saveAlertsWithCacheUseCase) handleSaveAlertsWithCache(ctx context.Context, alert domain.Alert) error {
	if err := uc.repo.Save(ctx, alert); err != nil {
		return err
	}

	key := extractCacheKeyFromAlert(alert)
	if alert.Status == domain.StatusResolved {
		return uc.cache.DeleteByKey(ctx, key)
	}

	if err := uc.cache.Save(ctx, key, alert); err != nil {
		return err
	}

	return nil

}
