package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	r "github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/application"
	ruleApplication "github.com/reneruprecht/alertbridge/backend/internal/application/rule"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alert"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alertmanager"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/postgres"
	rulePostgres "github.com/reneruprecht/alertbridge/backend/internal/infrastructure/postgres/rule"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/redis"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/rule"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

func main() {

	ctx := context.Background()

	connStr := "postgres://postgres:postgres@localhost:5432/alerts?sslmode=disable"
	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := postgres_db.New(conn)

	repo := postgres.NewAlertRepository(queries)
	ruleRepo := rulePostgres.NewRuleRepository(queries)

	opt, err := r.ParseURL("redis://:@localhost:6379/0")

	if err != nil {
		log.Fatal(err)
	}

	client := r.NewClient(opt)

	cache := redis.NewAlertCache(client)

	receiveAlertUsecase := application.NewReceiveAlertUseCase(repo, cache)
	findAlertsByInstanceUseCase := application.NewFindAlertsByInstanceUseCase(repo)
	listCachedAlertsUseCase := application.NewListActiveAlertsUseCase(cache)

	createRuleUseCase := ruleApplication.NewCreateRuleUseCase(ruleRepo)
	listRuleUseCase := ruleApplication.NewListRuleUseCase(ruleRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/alertmanager", alertmanager.HandleWebhook(receiveAlertUsecase))
	mux.HandleFunc("/api/v1/alerts/{instance}", alert.HandleFindAlertsByInstance(findAlertsByInstanceUseCase))
	mux.HandleFunc("/api/v1/alerts", alert.HandleListActiveAlerts(listCachedAlertsUseCase))

	mux.HandleFunc("POST /api/v1/rules", rule.HandleCreateRule(createRuleUseCase))
	mux.HandleFunc("GET /api/v1/rules", rule.HandleListRules(listRuleUseCase))

	log.Println("Server running on :8080")
	serverError := http.ListenAndServe(":8080", mux)

	if serverError != nil {
		log.Fatal(serverError)
	}
}
