//go:build integration
// +build integration

package postgres_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/reneruprecht/alertbridge/backend/internal/action/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/stretchr/testify/suite"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
)

type ActionRepositoryTestSuite struct {
	suite.Suite
	pgContainer *pg.PostgresContainer
	repo        *postgres.ActionRepository
	conn        *pgx.Conn
	ctx         context.Context
}

func (suite *ActionRepositoryTestSuite) SetupSuite() {
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

	suite.repo = postgres.NewActionRepository(queries)

}
func (suite *ActionRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		suite.T().Logf("error terminating postgres container: %s", err)
	}
}

func (suite *ActionRepositoryTestSuite) clearTables() {
	_, err := suite.conn.Exec(suite.ctx, "TRUNCATE actions CASCADE")
	suite.Require().NoError(err)
}

func (suite *ActionRepositoryTestSuite) SetupTest() {
	suite.clearTables()
}

func getSingleAction() domain.Action {

	id, _ := domain.NewActionID()
	name, _ := domain.NewActionName("webhook-sender")
	desc := domain.NewActionDescription("sends message")
	actionType, _ := domain.NewActionType("webhook")
	config, _ := domain.NewActionConfig(map[string]string{"url": "testurl"})

	action := domain.Action{ID: id, Name: name, Description: desc, Type: actionType, Config: config}
	return action
}

func (suite *ActionRepositoryTestSuite) TestActionRepository_Save() {
	action := getSingleAction()

	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	repoError := suite.repo.Save(ctx, action)
	suite.Require().NoError(repoError)

	var count int
	err := suite.conn.QueryRow(ctx, "SELECT count(*) FROM actions WHERE id=$1", action.ID).Scan(&count)

	suite.Require().NoError(err)
	suite.Assert().Equal(1, count)
}

func (suite *ActionRepositoryTestSuite) TestActionRepository_List() {

	action := getSingleAction()
	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	repoError := suite.repo.Save(ctx, action)
	suite.Require().NoError(repoError)

	actions, err := suite.repo.List(ctx)
	suite.Require().NoError(err)

	suite.Assert().Equal(1, len(actions))
}

func (suite *ActionRepositoryTestSuite) TestActionRepository_FindByID_Valid() {

	data := getSingleAction()
	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	repoError := suite.repo.Save(ctx, data)
	suite.Require().NoError(repoError)

	action, err := suite.repo.FindByID(ctx, data.ID)
	suite.Require().NoError(err)

	suite.Assert().Equal(data.ID, action.ID)
	suite.Assert().Equal(data.Name, action.Name)
	suite.Assert().Equal(data.Description, action.Description)
	suite.Assert().Equal(data.Type, action.Type)
	suite.Assert().Equal(data.Config, action.Config)
}

func (suite *ActionRepositoryTestSuite) TestActionRepository_FindByID_Error() {

	data := getSingleAction()
	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()

	_, err := suite.repo.FindByID(ctx, data.ID)
	suite.Require().Error(err)

}
func TestActionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ActionRepositoryTestSuite))
}
