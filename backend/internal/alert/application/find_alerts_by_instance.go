package application

import (
	"context"
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)


type FindAlertsByInstanceUseCaseInterface interface {
	Execute(ctx context.Context, instance string) ([]domain.Alert, error)
}

type FindAlertsByInstanceUseCase struct {
	repo AlertFindByInstanceRepository
}

func NewFindAlertsByInstanceUseCase(repo AlertFindByInstanceRepository) *FindAlertsByInstanceUseCase {
	return &FindAlertsByInstanceUseCase{repo: repo}
}

func (uc *FindAlertsByInstanceUseCase) Execute(ctx context.Context, instance string) ([]domain.Alert, error) {

	alerts, err := uc.repo.FindAlertsByInstance(ctx, instance)

	if err != nil {
		log.Printf("FindAlertsByInstanceUseCase error %v", err)
		return nil, err
	}
	return alerts, nil
}
