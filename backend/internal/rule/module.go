package rule

import (
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	httpCreateRule "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/create_rule"
	httpListRules "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/list_rules"
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

func (m *RuleModule) RegisterRuleRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/v1/rules", httpListRules.HandleListRules(m.ListRules))
	mux.HandleFunc("/api/v1/rules", httpCreateRule.HandleCreateRule(m.CreateRule))

}
