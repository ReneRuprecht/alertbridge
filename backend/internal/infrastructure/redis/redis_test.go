//go:build integration

package redis_test

import (
	"context"
	"testing"

	r "github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/application"
	"github.com/reneruprecht/alertbridge/backend/internal/domain"
	"github.com/reneruprecht/alertbridge/backend/internal/infrastructure/redis"
	"github.com/stretchr/testify/suite"
	tcr "github.com/testcontainers/testcontainers-go/modules/redis"
)

type AlertCacheTestSuite struct {
	suite.Suite
	redisContainer *tcr.RedisContainer
	cache          application.AlertCache
	client         *r.Client
	ctx            context.Context
}

func (suite *AlertCacheTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	redisContainer, err := tcr.Run(suite.ctx,
		"redis:8.2.3-alpine",
		tcr.WithLogLevel(tcr.LogLevelVerbose),
		tcr.WithSnapshotting(10, 1),
	)
	suite.Require().NoError(err)
	suite.redisContainer = redisContainer

	connStr, err := redisContainer.ConnectionString(suite.ctx)

	suite.Require().NoError(err)

	opt, err := r.ParseURL(connStr)
	suite.Require().NoError(err)

	client := r.NewClient(opt)

	suite.client = client

	suite.cache = redis.NewAlertCache(client)

}
func (suite *AlertCacheTestSuite) TearDownSuite() {
	if err := suite.redisContainer.Terminate(suite.ctx); err != nil {
		suite.T().Logf("error terminating redis container: %s", err)
	}
}

func (suite *AlertCacheTestSuite) TestAlertCache_Save() {

	fp, _ := domain.NewFingerprint("x123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")
	labels := make(map[string]string)
	labels["instance"] = "testinstance"
	labels["alertname"] = "testalert"

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt, Labels: labels}

	cacheError := suite.cache.Save(suite.ctx, alert)
	suite.Require().NoError(cacheError)

	res, err := suite.client.Get(suite.ctx, "alert:x123").Result()

	suite.Require().NoError(err)
	suite.Assert().NotEqual("", res)
	suite.Assert().Contains(res, "x123")
	suite.Assert().Contains(res, "firing")
	suite.Assert().Contains(res, "testinstance")
	suite.Assert().Contains(res, "testalert")
}

func (suite *AlertCacheTestSuite) TestAlertCache_ListAlerts() {

	fp, _ := domain.NewFingerprint("x123")
	status, _ := domain.NewStatus("firing")
	startsAt, _ := domain.NewTimestamp("2026-01-01T10:00:00Z")
	labels := make(map[string]string)
	labels["instance"] = "testinstance"
	labels["alertname"] = "testalert"

	alert := domain.Alert{Fingerprint: fp, Status: status, StartAt: startsAt, Labels: labels}

	cacheError := suite.cache.Save(suite.ctx, alert)
	suite.Require().NoError(cacheError)

	res, err := suite.cache.ListAlerts(suite.ctx)
	suite.Require().NoError(err)

	suite.Equal(1, len(res))
	suite.Equal("testinstance", res[0].Instance)
	suite.Equal("testalert", res[0].AlertName)
	suite.Equal("firing", res[0].Status)

}

func TestAlertRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AlertCacheTestSuite))
}
