package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/reneruprecht/alertbridge/backend/internal/application"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alert"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/alertmanager"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/postgres"
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

	receiveAlertUsecase := application.NewReceiveAlertUseCase(repo)
	findAlertsByInstanceUseCase := application.NewFindAlertsByInstanceUseCase(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/alertmanager", alertmanager.HandleWebhook(receiveAlertUsecase))
	mux.HandleFunc("/api/v1/alerts", alert.FindAlertsByInstance(findAlertsByInstanceUseCase))

	log.Println("Server running on :8080")
	serverError := http.ListenAndServe(":8080", mux)

	if serverError != nil {
		log.Fatal(serverError)
	}
}
