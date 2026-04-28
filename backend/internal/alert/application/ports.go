package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type AlertRepositoryWriter interface {
	Save(context context.Context, alert domain.Alert) error
}

type AlertRepositoryReader interface {
	ListByInstance(context context.Context, instance string) ([]domain.Alert, error)
}

type AlertCacheReader interface {
	List(context context.Context) ([]AlertCacheDto, error)
}

type AlertCacheWriter interface {
	Save(context context.Context,key string, alert domain.Alert) error
	DeleteByKey(context context.Context, key string) error
}
