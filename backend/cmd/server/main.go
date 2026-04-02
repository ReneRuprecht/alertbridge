package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/reneruprecht/alertbridge/backend/internal/application"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alertmanager"
	alertpostgres "github.com/reneruprecht/alertbridge/backend/internal/infrastructure/psql"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/psql"
)

func main() {

	ctx := context.Background()

    connStr := "postgres://postgres:postgres@localhost:5432/alerts?sslmode=disable"
	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

    queries:= psql.New(conn)

    repo := alertpostgres.NewAlertRepository(queries)

    receiveAlertUsecase := application.NewReceiveAlertUseCase(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/alerts", alertmanager.HandleWebhook(receiveAlertUsecase))

	log.Println("Server running on :8080")
	serverError := http.ListenAndServe(":8080", mux)

	if serverError != nil {
		log.Fatal(serverError)
	}
}
