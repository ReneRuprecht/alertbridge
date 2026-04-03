package application

import (
	"context"
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

type ReceiveAlertUsecase struct {
	repo domain.AlertRepository
}

func NewReceiveAlertUseCase(repo domain.AlertRepository) *ReceiveAlertUsecase {
	return &ReceiveAlertUsecase{repo: repo}
}

func (uc *ReceiveAlertUsecase) Execute(ctx context.Context, alerts []domain.Alert) error {
	for _, a := range alerts {

		if err := uc.repo.Save(ctx, a); err != nil {
			log.Printf("ReceiveAlertUsecase error %v", err)
			return err
		}
	}
	return nil
}
