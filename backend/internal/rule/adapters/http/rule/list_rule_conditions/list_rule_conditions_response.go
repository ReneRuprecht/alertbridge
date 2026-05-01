package listruleconditions

type ListRuleConditionsResponse struct {
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	ID       string `json:"id"`
	RuleID   string `json:"rule_id"`
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}
