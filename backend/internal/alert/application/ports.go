package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type AlertSaveRepository interface {
	Save(context context.Context, alert domain.Alert) error
}

type AlertFindByInstanceRepository interface {
	FindAlertsByInstance(context context.Context, instance string) ([]domain.Alert, error)
}

type AlertCacheReader interface {
	ListAlerts(context context.Context) ([]AlertCacheDto, error)
}

type AlertCacheWriter interface {
	Save(context context.Context, alert domain.Alert) error
	DeleteByKey(context context.Context, alert domain.Alert) error
}
