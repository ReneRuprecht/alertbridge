package alert

import (
	"net/http"

	r "github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/redis"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	httpAlertmanagerHandleWehook "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alertmanager"
	httpListAlertsByInstance "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_alerts_by_instance"
	httpListActiveAlerts "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_active_alerts"
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

func(m *AlertModule) RegisterAlertRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /api/v1/alertmanager", httpAlertmanagerHandleWehook.HandleWebhook(m.SaveAlertsWithCache))

	mux.HandleFunc("GET /api/v1/alerts/{instance}", httpListAlertsByInstance.HandleListAlertsByInstance(m.ListAlertsByInstance))
	mux.HandleFunc("GET /api/v1/alerts", httpListActiveAlerts.HandleListActiveAlerts(m.ListActiveAlerts))

}
