package application

import (
	"context"
)

type ListActiveAlertsUseCase interface {
	Execute(ctx context.Context) ([]AlertCacheDto, error)
}

type listActiveAlertsUseCase struct {
	cache AlertCacheReader
}

func NewListActiveAlertsUseCase(cache AlertCacheReader) *listActiveAlertsUseCase {
	return &listActiveAlertsUseCase{cache: cache}
}

func (uc *listActiveAlertsUseCase) Execute(ctx context.Context) ([]AlertCacheDto, error) {
	return uc.cache.List(ctx)
}
