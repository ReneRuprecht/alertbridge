package application

import (
	"context"
	"log"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type AlertProcessorUseCase interface {
	Execute(ctx context.Context, alertID domain.Fingerprint) error
}

type alertProcessorUseCase struct {
}

func NewAlertProcessorUseCase() *alertProcessorUseCase {
	return &alertProcessorUseCase{}
}

func (a *alertProcessorUseCase) Execute(ctx context.Context, alertID domain.Fingerprint) error {
	log.Printf("alert: %s", string(alertID))
	return nil
}
