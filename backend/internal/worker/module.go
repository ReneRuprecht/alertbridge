package worker

import (
	"github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/rabbitmq"
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

func NewAlertWorkerMoudle(rabbitWorkerConfig RabbitWorkerConfig) *AlertWorkerModule {

	processor := application.NewAlertProcessorUseCase()
	consumer := rabbitmq.NewAlertEventConsumer(rabbitWorkerConfig.Channel, rabbitWorkerConfig.Queue, processor)

	return &AlertWorkerModule{Consumer: *consumer}
}
