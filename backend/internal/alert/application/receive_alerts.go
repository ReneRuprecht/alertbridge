package application

import (
	"context"
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type ReceiveAlertUsecaseInterface interface {
	Execute(ctx context.Context, alerts []domain.Alert) error
}

type ReceiveAlertUsecase struct {
	repo  AlertSaveRepository
	cache AlertCacheWriter
}

func NewReceiveAlertUseCase(repo AlertSaveRepository, cache AlertCacheWriter) *ReceiveAlertUsecase {
	return &ReceiveAlertUsecase{repo: repo, cache: cache}
}

func (uc *ReceiveAlertUsecase) Execute(ctx context.Context, alerts []domain.Alert) error {
	for _, alert := range alerts {

		if err := uc.repo.Save(ctx, alert); err != nil {
			log.Printf("ReceiveAlertUsecase repo error %v", err)
			return err
		}

		if alert.Status == "resolved" {
			err := uc.cache.DeleteByKey(ctx, alert)

			if err != nil {
				return err
			}
			continue
		}

		if err := uc.cache.Save(ctx, alert); err != nil {
			log.Printf("ReceiveAlertUsecase cache error %v", err)
			return err
		}
	}
	return nil
}
