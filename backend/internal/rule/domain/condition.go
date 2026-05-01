package domain

type Condition struct {
	ID       ConditionID
	RuleID   RuleId
	Name     ConditionName
	Operator ConditionOperator
	Value    ConditionValue
	Field    ConditionField
}
