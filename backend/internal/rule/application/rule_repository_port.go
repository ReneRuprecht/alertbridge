package application

import (
	"context"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

type RuleRepository interface {
	Save(context context.Context, rule domain.Rule) error
	List(context context.Context) ([]domain.Rule, error)
}
