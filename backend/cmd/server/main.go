package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	r "github.com/redis/go-redis/v9"
	httpListActiveAlerts "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_active_alerts"
	httpListAlertsByInstance "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alert/list_alerts_by_instance"
	alertHttpAlertmanager "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/http/alertmanager"
	alertPostgres "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/postgres"
	alertRedis "github.com/reneruprecht/alertbridge/backend/internal/alert/adapters/redis"
	alertApplication "github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	httpCreateRule "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/create_rule"
	httpListRules "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/list_rules"
	rulePostgres "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/postgres"
	ruleApplication "github.com/reneruprecht/alertbridge/backend/internal/rule/application"
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

	repo := alertPostgres.NewAlertRepository(queries)
	ruleRepo := rulePostgres.NewRuleRepository(queries)

	opt, err := r.ParseURL("redis://:@localhost:6379/0")

	if err != nil {
		log.Fatal(err)
	}

	client := r.NewClient(opt)

	cache := alertRedis.NewAlertCache(client)

	saveAlertsWithCacheUseCase := alertApplication.NewSaveAlertsWithCacheUseCase(repo, cache)
	listAlertsByInstance := alertApplication.NewListAlertsByInstanceUseCase(repo)
	listActiveAlertsUseCase := alertApplication.NewListActiveAlertsUseCase(cache)

	createRuleUseCase := ruleApplication.NewCreateRuleUseCase(ruleRepo)
	listRuleUseCase := ruleApplication.NewListRuleUseCase(ruleRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/alertmanager", alertHttpAlertmanager.HandleWebhook(saveAlertsWithCacheUseCase))

	mux.HandleFunc("GET /api/v1/alerts/{instance}", httpListAlertsByInstance.HandleListAlertsByInstance(listAlertsByInstance))
	mux.HandleFunc("GET /api/v1/alerts", httpListActiveAlerts.HandleListActiveAlerts(listActiveAlertsUseCase))

	mux.HandleFunc("POST /api/v1/rules", httpCreateRule.HandleCreateRule(createRuleUseCase))
	mux.HandleFunc("GET /api/v1/rules", httpListRules.HandleListRules(listRuleUseCase))

	log.Println("Server running on :8080")
	serverError := http.ListenAndServe(":8080", mux)

	if serverError != nil {
		log.Fatal(serverError)
	}
}
