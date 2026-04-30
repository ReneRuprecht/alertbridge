package alert

import (
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
	r "github.com/redis/go-redis/v9"
	httpListActiveAlerts "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_active_alerts"
	httpListAlertsByInstance "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_alerts_by_instance"
	httpAlertmanagerHandleWehook "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alertmanager"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/rabbitmq"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/redis"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

type RabbitConfig struct {
	Channel *amqp.Channel
	Queue   string
}

type AlertModule struct {
	SaveAlertsWithCache  application.SaveAlertsWithCacheUseCase
	ListAlertsByInstance application.ListAlertsByInstanceUseCase
	ListActiveAlerts     application.ListActiveAlertsUseCase
	PublishAlert         application.PublishAlertUsecase
}

func NewAlertModule(queries *postgres_db.Queries, client *r.Client, rabbitConfig *RabbitConfig) (*AlertModule, error) {
	repo := postgres.NewAlertRepository(queries)
	cache := redis.NewAlertCache(client)

	publisher := rabbitmq.NewAlertEventPublisher(rabbitConfig.Channel, rabbitConfig.Queue)
	err := publisher.Init()

	if err != nil {
		return nil, err
	}

	return &AlertModule{
		SaveAlertsWithCache:  application.NewSaveAlertsWithCacheUseCase(repo, cache),
		ListAlertsByInstance: application.NewListAlertsByInstanceUseCase(repo),
		ListActiveAlerts:     application.NewListActiveAlertsUseCase(cache),
		PublishAlert:         application.NewAlertEventPublisherUseCase(publisher),
	}, nil
}

func (m *AlertModule) RegisterAlertRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /api/v1/alertmanager", httpAlertmanagerHandleWehook.HandleWebhook(m.SaveAlertsWithCache, m.PublishAlert))

	mux.HandleFunc("GET /api/v1/alerts/{instance}", httpListAlertsByInstance.HandleListAlertsByInstance(m.ListAlertsByInstance))
	mux.HandleFunc("GET /api/v1/alerts", httpListActiveAlerts.HandleListActiveAlerts(m.ListActiveAlerts))

}
