package application

import (
	"context"
	"errors"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type RuleRepository interface {
	Save(context context.Context, rule domain.Rule) error
	List(context context.Context) ([]domain.Rule, error)
}

var ErrorRuleNotFound = errors.New("rule not found")

type RuleFinder interface {
	FindByID(context context.Context, ruleID domain.RuleId) (domain.Rule, error)
}
