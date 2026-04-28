package rule

import (
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
)

type RuleModule struct {
	CreateRule application.CreateRuleUseCase
	ListRules  application.ListRulesUseCase
}

func NewRuleModule(queries *postgres_db.Queries) *RuleModule {
	repo := postgres.NewRuleRepository(queries)

	return &RuleModule{
		CreateRule: application.NewCreateRuleUseCase(repo),
		ListRules:  application.NewListRuleUseCase(repo),
	}
}
