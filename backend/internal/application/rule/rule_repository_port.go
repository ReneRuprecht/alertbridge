package rule

import (
	"context"

	domain "github.com/reneruprecht/alertbridge/backend/internal/domain/rule"
)

type RuleRepository interface {
	Save(context context.Context, rule domain.Rule) error
	List(context context.Context) ([]domain.Rule, error)
}
