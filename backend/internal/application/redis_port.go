package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/domain"
)

type AlertCache interface {
	ListAlerts(context context.Context) ([]AlertCacheDto, error)
	Save(context context.Context, alert domain.Alert) error
	DeleteByKey(context context.Context, alert domain.Alert) error
}
