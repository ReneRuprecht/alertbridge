package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	r "github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/action"
	"github.com/reneruprecht/alertbridge/backend/internal/alert"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/rabbitmq"
	"github.com/reneruprecht/alertbridge/backend/internal/rule"
)

type config struct {
	PostgresConnStr  string
	RedisConnStr     string
	RabbitMQConnStr  string
	Port             string
	PublisherEnabled bool
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

	serverPort := os.Getenv("AB_SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	publisherEnabled := false
	if os.Getenv("AB_PUBLISHER_ENABLED") == "1" {
		publisherEnabled = true
	}

	return config{PostgresConnStr: postgresConnStr, RabbitMQConnStr: rabbitMQConnStr, RedisConnStr: redisConnStr, Port: serverPort, PublisherEnabled: publisherEnabled}

}

func setupRedisClient(cfg config) (*r.Client, error) {
	opt, err := r.ParseURL(cfg.RedisConnStr)

	if err != nil {
		return nil, err
	}

	client := r.NewClient(opt)

	return client, nil
}

func startServer(mux *http.ServeMux, cfg config) {

	serverPort := os.Getenv("AB_SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	log.Printf("Server running on :%s\n", serverPort)
	serverError := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), mux)

	if serverError != nil {
		log.Fatal(serverError)
	}
}

func main() {

	ctx := context.Background()

	cfg := loadConfig()

	postgresConn, err := pgx.Connect(ctx, cfg.PostgresConnStr)

	if err != nil {
		log.Fatal(err)
	}

	defer postgresConn.Close(ctx)

	queries := postgres_db.New(postgresConn)

	redisClient, err := setupRedisClient(cfg)

	if err != nil {
		log.Fatal(err)
	}

	rabbit, err := rabbitmq.NewRabbit(cfg.RabbitMQConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()

	alertCh, err := rabbit.Channel()
	if err != nil {
		log.Fatal(err)
	}

	alertModule, err := alert.NewAlertModule(queries, redisClient, &alert.RabbitConfig{Channel: alertCh, Queue: "alerts", Enabled: cfg.PublisherEnabled})
	if err != nil {
		log.Fatal(err)
	}

	ruleModule := rule.NewRuleModule(queries)
	actionModule := action.NewActionModule(queries)

	mux := http.NewServeMux()

	alertModule.RegisterAlertRoutes(mux)
	ruleModule.RegisterRuleRoutes(mux)
	actionModule.RegisterAlertRoutes(mux)

	startServer(mux, cfg)

}
