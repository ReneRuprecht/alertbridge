package listconditionsbyruleid

type ListConditionsByRuleIDResponse struct {
	Rule       Rule        `json:"rule"`
	Conditions []Condition `json:"conditions"`
}

type Rule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Enabled     bool   `json:"enabled"`
}
type Condition struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}
