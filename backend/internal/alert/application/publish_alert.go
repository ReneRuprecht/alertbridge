package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type PublishAlertUsecase interface {
	Execute(ctx context.Context, alerts []domain.Alert) error
}

type publishAlertUsecase struct {
	publisher AlertEventPublisher
}

func NewAlertEventPublisherUseCase(publisher AlertEventPublisher) *publishAlertUsecase {
	return &publishAlertUsecase{publisher: publisher}
}

func (uc *publishAlertUsecase) Execute(ctx context.Context, alerts []domain.Alert) error {
	for _, alert := range alerts {

		if err := uc.publisher.Publish(ctx, alert.Fingerprint); err != nil {
			return err
		}
	}
	return nil
}

type fakePublishAlertUsecase struct {
}

func NewFakeAlertEventPublisherUseCase() *fakePublishAlertUsecase {
	return &fakePublishAlertUsecase{}
}

func (uc *fakePublishAlertUsecase) Execute(ctx context.Context, alerts []domain.Alert) error {
	return nil
}
