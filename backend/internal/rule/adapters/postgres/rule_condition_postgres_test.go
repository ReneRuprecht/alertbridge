package postgres_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/suite"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
)

type RuleConditionRepositoryTestSuite struct {
	suite.Suite
	pgContainer *pg.PostgresContainer
	repo        application.RuleConditionRepository
	ruleRepo    application.RuleRepository
	conn        *pgx.Conn
	ctx         context.Context
}

func (suite *RuleConditionRepositoryTestSuite) SetupSuite() {
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

	suite.repo = postgres.NewRuleConditionRepository(queries)
	suite.ruleRepo = postgres.NewRuleRepository(queries)

	rule := getRule()
	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	suite.ruleRepo.Save(ctx, rule)

}
func (suite *RuleConditionRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		suite.T().Logf("error terminating postgres container: %s", err)
	}
}

func (suite *RuleConditionRepositoryTestSuite) clearTables() {
	_, err := suite.conn.Exec(suite.ctx, "TRUNCATE rule_conditions CASCADE")
	suite.Require().NoError(err)
}

func (suite *RuleConditionRepositoryTestSuite) SetupTest() {
	suite.clearTables()
}

func getRule() domain.Rule {

	id, _ := uuid.Parse("1018f6b7e-2c4a-7f3a-9c2d-1a2b3c4d5e6f")
	name, _ := domain.NewRuleName("testName")
	description := "testDesc"
	priority, _ := domain.NewRulePriority(100)
	enabled := true

	rule := domain.Rule{
		ID:          domain.RuleId(id),
		Name:        name,
		Description: description,
		Priority:    priority,
		Enabled:     enabled,
	}
	return rule
}

func getCondition() domain.Condition {
	id, _ := domain.NewConditionID()
	name, _ := domain.NewConditionName("status critical")
	operator, _ := domain.NewConditionOperator("equals")
	field, _ := domain.NewConditionField("status")
	value, _ := domain.NewConditionValue("critical")

	ruleIDParsed, _ := uuid.Parse("1018f6b7e-2c4a-7f3a-9c2d-1a2b3c4d5e6f")
	ruleID := domain.RuleId(ruleIDParsed)

	return domain.Condition{
		ID:       id,
		RuleID:   ruleID,
		Name:     name,
		Operator: operator,
		Field:    field,
		Value:    value,
	}
}

func (suite *RuleConditionRepositoryTestSuite) TestRuleConditionRepository_SaveOne() {

	condition := getCondition()

	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	repoError := suite.repo.Save(ctx, condition)
	suite.Require().NoError(repoError)

	var count int
	err := suite.conn.QueryRow(suite.ctx, "SELECT count(*) FROM rule_conditions WHERE id=$1", condition.ID).Scan(&count)

	suite.Require().NoError(err)
	suite.Assert().Equal(1, count)
}

func (suite *RuleConditionRepositoryTestSuite) TestRuleConditionRepository_List() {

	condition := getCondition()

	ctx, cancel := context.WithTimeout(suite.ctx, 10*time.Second)
	defer cancel()
	repoError := suite.repo.Save(ctx, condition)
	suite.Require().NoError(repoError)

	conditions, err := suite.repo.List(ctx)

	suite.Require().NoError(err)
	suite.Assert().Equal(1, len(conditions))

}
func TestRuleConditionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RuleConditionRepositoryTestSuite))
}
