package application

import (
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

type ReceiveAlertUsecase struct {
	repo domain.AlertRepository
}

func NewReceiveAlertUseCase(repo domain.AlertRepository) *ReceiveAlertUsecase {
	return &ReceiveAlertUsecase{repo: repo}
}

func (uc *ReceiveAlertUsecase) Execute(alerts []domain.Alert) error {
	for _, a := range alerts {

		if err := uc.repo.Save(a); err != nil {
            log.Printf("ReceiveAlertUsecase error %v", err)
			return err
		}
	}
	return nil
}
