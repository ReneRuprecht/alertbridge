package application

import (
	"context"
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

type ReceiveAlertUsecase struct {
	repo  domain.AlertRepository
	cache AlertCache
}

func NewReceiveAlertUseCase(repo domain.AlertRepository, cache AlertCache) *ReceiveAlertUsecase {
	return &ReceiveAlertUsecase{repo: repo, cache: cache}
}

func (uc *ReceiveAlertUsecase) Execute(ctx context.Context, alerts []domain.Alert) error {
	for _, alert := range alerts {

		if err := uc.repo.Save(ctx, alert); err != nil {
			log.Printf("ReceiveAlertUsecase repo error %v", err)
			return err
		}

		if alert.Status == "resolved" {
			continue
		}

		if err := uc.cache.Save(ctx, alert); err != nil {
			log.Printf("ReceiveAlertUsecase cache error %v", err)
			return err
		}
	}
	return nil
}
