package alert

import (
	r "github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/redis"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

type AlertModule struct {
	SaveAlertsWithCache  application.SaveAlertsWithCacheUseCase
	ListAlertsByInstance application.ListAlertsByInstanceUseCase
	ListActiveAlerts     application.ListActiveAlertsUseCase
}

func NewAlertModule(queries *postgres_db.Queries, client *r.Client) *AlertModule {
	repo := postgres.NewAlertRepository(queries)
	cache := redis.NewAlertCache(client)

	return &AlertModule{
		SaveAlertsWithCache:  application.NewSaveAlertsWithCacheUseCase(repo, cache),
		ListAlertsByInstance: application.NewListAlertsByInstanceUseCase(repo),
		ListActiveAlerts:     application.NewListActiveAlertsUseCase(cache),
	}
}
