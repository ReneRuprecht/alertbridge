package application

import (
	"context"
	"fmt"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type FindAlertByFingerprintUseCase interface {
	Execute(ctx context.Context, alertFingerprint domain.Fingerprint) (AlertCacheDto, error)
}

type findAlertByFingerprintUseCase struct {
	cache AlertCacheFinder
}

func NewFindAlertByFingerprintUseCase(cache AlertCacheFinder) *findAlertByFingerprintUseCase {
	return &findAlertByFingerprintUseCase{cache: cache}
}

func (uc *findAlertByFingerprintUseCase) Execute(ctx context.Context, alertFingerprint domain.Fingerprint) (AlertCacheDto, error) {
	key := fmt.Sprintf("alert:%s", alertFingerprint)

	return uc.cache.FindByKey(ctx, key)
}
