package rule

type Rule struct {
	ID          RuleId
	Name        RuleName
	Description string
	Priority    RulePriority
	Enabled     bool
}
