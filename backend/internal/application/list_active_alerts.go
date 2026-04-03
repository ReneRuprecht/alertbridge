package application

import (
	"context"
	"log"
)

type ListActiveAlertsUseCaseInterface interface {
	Execute([]AlertCacheDto, error)
}

type ListActiveAlertsUseCase struct {
	cache AlertCache
}

func NewListActiveAlertsUseCase(cache AlertCache) *ListActiveAlertsUseCase {
	return &ListActiveAlertsUseCase{cache: cache}
}

func (uc *ListActiveAlertsUseCase) Execute(ctx context.Context) ([]AlertCacheDto, error) {

	alerts, err := uc.cache.ListAlerts(ctx)

	if err != nil {
		log.Printf("FindAlertsByInstanceUseCase error %v", err)
		return nil, err
	}
	return alerts, nil
}
