//go:build integration
// +build integration

package rule_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5"
	application "github.com/reneruprecht/alertbridge/backend/internal/application/rule"
	domain "github.com/reneruprecht/alertbridge/backend/internal/domain/rule"
	postgres "github.com/reneruprecht/alertbridge/backend/internal/infrastructure/postgres/rule"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/stretchr/testify/suite"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
)

type RuleRepositoryTestSuite struct {
	suite.Suite
	pgContainer *pg.PostgresContainer
	repo        application.RuleRepository
	conn        *pgx.Conn
	ctx         context.Context
}

func (suite *RuleRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	dbName := "alerts"
	dbUser := "postgres"
	dbPassword := "postgres"

	postgresContainer, err := pg.Run(suite.ctx,
		"postgres:18-alpine",
		pg.WithInitScripts(filepath.Join("../../../platform/postgres_db/", "schema.sql")),
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

	suite.repo = postgres.NewRuleRepository(queries)

}
func (suite *RuleRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		suite.T().Logf("error terminating postgres container: %s", err)
	}
}

func (suite *RuleRepositoryTestSuite) clearTables() {
	_, err := suite.conn.Exec(suite.ctx, "TRUNCATE rules CASCADE")
	suite.Require().NoError(err)
}

func (suite *RuleRepositoryTestSuite) SetupTest() {
	suite.clearTables()
}

func (suite *RuleRepositoryTestSuite) TestRuleRepository_SaveOne() {

	id := domain.NewRuleId()
	name, _ := domain.NewRuleName("testName")
	description := "testDesc"
	priority,_ := domain.NewRulePriority(100)
	enabled := true

	rule := domain.Rule{
		ID:          id,
		Name:        name,
		Description: description,
		Priority:    priority,
		Enabled:     enabled,
	}

	repoError := suite.repo.Save(suite.ctx, rule)
	suite.Require().NoError(repoError)

	var count int
	err := suite.conn.QueryRow(suite.ctx, "SELECT count(*) FROM rules WHERE id=$1", id).Scan(&count)

	suite.Require().NoError(err)
	suite.Assert().Equal(1, count)
}

func (suite *RuleRepositoryTestSuite) TestRuleRepository_List() {

	id := domain.NewRuleId()
	name, _ := domain.NewRuleName("testName")
	description := "testDesc"
	priority,_ := domain.NewRulePriority(100)
	enabled := true

	rule := domain.Rule{
		ID:          id,
		Name:        name,
		Description: description,
		Priority:    priority,
		Enabled:     enabled,
	}

	repoError := suite.repo.Save(suite.ctx, rule)
	suite.Require().NoError(repoError)

	rules, err := suite.repo.List(suite.ctx)

	suite.Require().NoError(err)
	suite.Assert().Equal(1, len(rules))
}
func TestRuleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RuleRepositoryTestSuite))
}
