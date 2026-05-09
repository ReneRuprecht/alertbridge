package application

import (
	"context"
	"fmt"

	alertApp "github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	alertDomain "github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
	ruleApp "github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	ruleDomain "github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type AlertProcessorUseCase interface {
	Execute(ctx context.Context, alertID alertDomain.Fingerprint) ([]ruleDomain.Rule, error)
}

type alertProcessorUseCase struct {
	alertFinder   alertApp.AlertCacheFinder
	ruleRepo      ruleApp.RuleRepository
	conditionRepo ruleApp.RuleConditionRepository
	ruleMatcher   ruleApp.MatchRuleConditionUseCase
}

func NewAlertProcessorUseCase(alertFinger alertApp.AlertCacheFinder, ruleRepo ruleApp.RuleRepository, conditionRepo ruleApp.RuleConditionRepository, ruleMatcher ruleApp.MatchRuleConditionUseCase) *alertProcessorUseCase {
	return &alertProcessorUseCase{alertFinder: alertFinger, ruleRepo: ruleRepo, conditionRepo: conditionRepo, ruleMatcher: ruleMatcher}
}

func (a *alertProcessorUseCase) Execute(ctx context.Context, alertFingerprint alertDomain.Fingerprint) ([]ruleDomain.Rule, error) {
	key := fmt.Sprintf("alert:%s", string(alertFingerprint))
	alert, err := a.alertFinder.FindByKey(ctx, key)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rules, err := a.ruleRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	conditions, err := a.conditionRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	groups := groupConditionsByRuleID(conditions)

	fmt.Printf("loaded %d groups\n", len(groups))

	matched := []ruleDomain.Rule{}

	for _, r := range rules {
		ruleConditions, ok := groups[r.ID]

		if !ok {
			fmt.Printf("no conditions found for rule %s\n", string(r.Name))
			fmt.Println("------------")
			continue
		}

		ruleMatches := true

		for _, c := range ruleConditions {
			if !a.ruleMatcher.Execute(alert, c) {
				fmt.Println("condition not matched")
				fmt.Println(c)
				fmt.Println("------------")
				ruleMatches = false
				break
			}
		}
		if ruleMatches {
			fmt.Printf("found matching rule %s", string(r.Name))
			fmt.Println(r)
			fmt.Println("------")
			matched = append(matched, r)
		}
	}

	return matched, nil
}

func groupConditionsByRuleID(conditions []ruleDomain.Condition) map[ruleDomain.RuleId][]ruleDomain.Condition {
	result := make(map[ruleDomain.RuleId][]ruleDomain.Condition)

	for _, c := range conditions {
		result[c.RuleID] = append(result[c.RuleID], c)
	}

	return result
}
