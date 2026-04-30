package main

import (
	"context"
	"log"
	"os"

	"github.com/reneruprecht/alertbridge/backend/internal/platform/rabbitmq"
	"github.com/reneruprecht/alertbridge/backend/internal/worker"
)

type config struct {
	RabbitMQConnStr string
	Port            string
}

func loadConfig() config {

	rabbitMQConnStr := os.Getenv("AB_RABBITMQ_CONNSTR")
	if rabbitMQConnStr == "" {
		rabbitMQConnStr = "amqp://rabbit:rabbit@localhost:5672/"
	}

	return config{RabbitMQConnStr: rabbitMQConnStr}

}

func main() {

	cfg := loadConfig()

	ctx := context.Background()

	rabbit, err := rabbitmq.NewRabbit(cfg.RabbitMQConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()

	alertChannel, err := rabbit.Channel()
	if err != nil {
		log.Fatal(err)
	}

	alertCfg := &worker.RabbitWorkerConfig{Channel: alertChannel, Queue: "alerts"}

	alertWorkerModule := worker.NewAlertWorkerMoudle(*alertCfg)

	alertWorkerModule.Consumer.Consume(ctx)

}
