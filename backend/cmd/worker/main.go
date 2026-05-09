package main

import (
	"context"
	"log"
	"os"

	r "github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/rabbitmq"
	"github.com/reneruprecht/alertbridge/backend/internal/worker"
)

type config struct {
	PostgresConnStr string
	RedisConnStr    string
	RabbitMQConnStr string
}

func loadConfig() config {

	postgresConnStr := os.Getenv("AB_POSTGRES_CONNSTR")
	if postgresConnStr == "" {
		postgresConnStr = "postgres://postgres:postgres@localhost:5432/alerts?sslmode=disable"
	}

	redisConnStr := os.Getenv("AB_REDIS_CONNSTR")
	if redisConnStr == "" {
		redisConnStr = "redis://:@localhost:6379/0"
	}

	rabbitMQConnStr := os.Getenv("AB_RABBITMQ_CONNSTR")
	if rabbitMQConnStr == "" {
		rabbitMQConnStr = "amqp://rabbit:rabbit@localhost:5672/"
	}

	return config{PostgresConnStr: postgresConnStr, RabbitMQConnStr: rabbitMQConnStr, RedisConnStr: redisConnStr}

}

func setupRedisClient(cfg config) (*r.Client, error) {
	opt, err := r.ParseURL(cfg.RedisConnStr)

	if err != nil {
		return nil, err
	}

	client := r.NewClient(opt)

	return client, nil
}

func main() {

	cfg := loadConfig()

	ctx := context.Background()

	postgresConn, err := pgx.Connect(ctx, cfg.PostgresConnStr)

	if err != nil {
		log.Fatal(err)
	}

	defer postgresConn.Close(ctx)

	queries := postgres_db.New(postgresConn)

	rabbit, err := rabbitmq.NewRabbit(cfg.RabbitMQConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()

	redisClient, err := setupRedisClient(cfg)

	if err != nil {
		log.Fatal(err)
	}

	alertChannel, err := rabbit.Channel()
	if err != nil {
		log.Fatal(err)
	}

	alertCfg := &worker.RabbitWorkerConfig{Channel: alertChannel, Queue: "alerts"}

	alertWorkerModule := worker.NewAlertWorkerMoudle(*alertCfg, redisClient, queries)

	alertWorkerModule.Consumer.Consume(ctx)

}
