package rule

import (
	"net/http"

	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
	httpAddRuleCondition "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/add_rule_condition"
	httpCreateRule "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/create_rule"
	httpListConditionsByRuleID "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/list_conditions_by_rule_id"
	httpListRuleConditions "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/list_rule_conditions"
	httpListRules "github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/http/rule/list_rules"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
)

type RuleModule struct {
	CreateRule             application.CreateRuleUseCase
	ListRules              application.ListRulesUseCase
	AddRuleCondition       application.AddRuleConditionUseCase
	ListRuleConditions     application.ListRuleConditionsUseCase
	ListConditionsByRuleID application.ListConditionsByRuleIDUseCase
}

func NewRuleModule(queries *postgres_db.Queries) *RuleModule {
	ruleRepo := postgres.NewRuleRepository(queries)
	conditionRepo := postgres.NewRuleConditionRepository(queries)

	return &RuleModule{
		CreateRule:             application.NewCreateRuleUseCase(ruleRepo),
		ListRules:              application.NewListRuleUseCase(ruleRepo),
		AddRuleCondition:       application.NewAddRuleConditionUseCase(ruleRepo, conditionRepo),
		ListRuleConditions:     application.NewListRuleConditionsUseCase(conditionRepo),
		ListConditionsByRuleID: application.NewListConditionsByRuleIDUseCase(ruleRepo, conditionRepo),
	}
}

func (m *RuleModule) RegisterRuleRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/v1/rules", httpListRules.HandleListRules(m.ListRules))
	mux.HandleFunc("/api/v1/rules", httpCreateRule.HandleCreateRule(m.CreateRule))
	mux.HandleFunc("/api/v1/rules/conditions", httpAddRuleCondition.HandleAddRuleCondition(m.AddRuleCondition))
	mux.HandleFunc("GET /api/v1/rules/conditions", httpListRuleConditions.HandleListRuleConditions(m.ListRuleConditions))
	mux.HandleFunc("/api/v1/rules/{id}/conditions", httpListConditionsByRuleID.HandleListConditionsByRuleID(m.ListConditionsByRuleID))

}
