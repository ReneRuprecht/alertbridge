// +build integration

package postgres_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/stretchr/testify/suite"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
)

type AlertRepositoryTestSuite struct {
	suite.Suite
	pgContainer *pg.PostgresContainer
	repo        domain.AlertRepository
	conn        *pgx.Conn
	ctx         context.Context
}

func (suite *AlertRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	dbName := "alerts"
	dbUser := "postgres"
	dbPassword := "postgres"

	postgresContainer, err := pg.Run(suite.ctx,
		"postgres:18-alpine",
		pg.WithInitScripts(filepath.Join("../../platform/postgres_db/", "schema.sql")),
		pg.WithDatabase(dbName),
		pg.WithUsername(dbUser),
		pg.WithPassword(dbPassword),
		pg.BasicWaitStrategies(),
	)

	suite.Require().NoError(err)

	suite.pgContainer = postgresContainer
	connStr, err := postgresContainer.ConnectionString(suite.ctx, "sslmode=disable")

	suite.Require().NoError(err)

	conn, err := pgx.Connect(suite.ctx, connStr)

	suite.Require().NoError(err)

	suite.conn = conn

	queries := postgres_db.New(conn)

	suite.repo = postgres.NewAlertRepository(queries)

}
func (suite *AlertRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		suite.T().Logf("error terminating postgres container: %s", err)
	}
}

func (suite *AlertRepositoryTestSuite) clearTables() {
	_, err := suite.conn.Exec(suite.ctx, "TRUNCATE alerts CASCADE")
	suite.Require().NoError(err)
}

func (suite *AlertRepositoryTestSuite) SetupTest() {
	suite.clearTables()
}

func (suite *AlertRepositoryTestSuite) TestAlertRepository_SaveOne() {

	fp, _ := domain.NewFingerprint("x123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt}

	repoError := suite.repo.Save(suite.ctx, alert)
	suite.Require().NoError(repoError)

	var count int
	err := suite.conn.QueryRow(suite.ctx, "SELECT count(*) FROM alerts WHERE fingerprint=$1", "x123").Scan(&count)

	suite.Require().NoError(err)
	suite.Assert().Equal(1, count)
}

func (suite *AlertRepositoryTestSuite) TestAlertRepository_SaveMany() {

	fp1, _ := domain.NewFingerprint("a123")
	status1, _ := domain.NewStatus("firing")
	startsAt1, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")

	fp2, _ := domain.NewFingerprint("b123")
	status2, _ := domain.NewStatus("resolved")
	startsAt2, _ := domain.NewTimestamp("2025-01-01T10:00:00Z")

	alerts := []domain.Alert{
		{
			Fingerprint: fp1,
			Status:      status1,
			StartAt:     startsAt1,
		},
		{
			Fingerprint: fp2,
			Status:      status2,
			StartAt:     startsAt2,
		},
	}

	for _, alert := range alerts {

		repoError := suite.repo.Save(suite.ctx, alert)
		suite.Assert().NoError(repoError)
	}

	var count int
	err := suite.conn.QueryRow(suite.ctx, "SELECT count(*) FROM alerts").Scan(&count)

	suite.Require().NoError(err)
	suite.Assert().Equal(2, count)

}

func (suite *AlertRepositoryTestSuite) TestAlertRepository_FindAlertsByInstance() {

	fp, _ := domain.NewFingerprint("x123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")
	label := make(map[string]string)
	label["instance"] = "testinstance"

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt, Labels: label}

	repoError := suite.repo.Save(suite.ctx, alert)
	suite.Require().NoError(repoError)

	alerts, err := suite.repo.FindAlertsByInstance(suite.ctx, "testinstance")

	suite.Require().NoError(err)
	suite.Assert().Equal(1, len(alerts))
}

func TestAlertRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AlertRepositoryTestSuite))
}
