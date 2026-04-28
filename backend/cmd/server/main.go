package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	r "github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/alert"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/rule"
)

type config struct {
	PostgresConnStr string
	RedisConnStr    string
	Port            string
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
	serverPort := os.Getenv("AB_SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	return config{PostgresConnStr: postgresConnStr, RedisConnStr: redisConnStr, Port: serverPort}

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

	alertModule := alert.NewAlertModule(queries, redisClient)
	ruleModule := rule.NewRuleModule(queries)

	mux := http.NewServeMux()

	alertModule.RegisterAlertRoutes(mux)
	ruleModule.RegisterRuleRoutes(mux)

	startServer(mux, cfg)

}
