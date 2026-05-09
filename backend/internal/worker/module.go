package worker

import (
	r "github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/rabbitmq"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/redis"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	rulePg "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/postgres"
	ruleApplication "github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/worker/application"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitWorkerConfig struct {
	Channel *amqp.Channel
	Queue   string
}

type AlertWorkerModule struct {
	Consumer rabbitmq.Consumer
}

func NewAlertWorkerMoudle(rabbitWorkerConfig RabbitWorkerConfig, client *r.Client, queries *postgres_db.Queries) *AlertWorkerModule {
	alertfinder := redis.NewAlertCache(client)

	matcher := ruleApplication.NewMatchRuleConditionUseCase()

	ruleRepo := rulePg.NewRuleRepository(queries)
	ruleConditionRepo := rulePg.NewRuleConditionRepository(queries)

	processor := application.NewAlertProcessorUseCase(alertfinder, ruleRepo, ruleConditionRepo, matcher)
	consumer := rabbitmq.NewAlertEventConsumer(rabbitWorkerConfig.Channel, rabbitWorkerConfig.Queue, processor)

	return &AlertWorkerModule{Consumer: *consumer}
}
