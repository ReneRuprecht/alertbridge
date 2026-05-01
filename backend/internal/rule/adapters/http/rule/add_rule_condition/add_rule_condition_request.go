package addrulecondition

type AddRuleConditionRequest struct {
	RuleID   string `json:"ruleID"`
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}
