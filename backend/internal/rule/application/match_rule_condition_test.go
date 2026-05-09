package application_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	alertApplication "github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"github.com/stretchr/testify/assert"
)

func getCondition() domain.Condition {
	id, _ := domain.NewConditionID()
	name, _ := domain.NewConditionName("status critical")
	operator, _ := domain.NewConditionOperator("equals")
	field, _ := domain.NewConditionField("status")
	value, _ := domain.NewConditionValue("firing")

	ruleIDParsed, _ := uuid.Parse("1018f6b7e-2c4a-7f3a-9c2d-1a2b3c4d5e6f")
	ruleID := domain.RuleId(ruleIDParsed)

	return domain.Condition{
		ID:       id,
		RuleID:   ruleID,
		Name:     name,
		Operator: operator,
		Field:    field,
		Value:    value,
	}
}

func getAlert() alertApplication.AlertCacheDto {

	fp1 := "a123"
	status1 := "firing"
	instance1 := "db01"
	job1 := "db-exporter"
	alertName1 := "db down"
	severity1 := "critical"
	startsAt1, _ := time.Parse(time.RFC3339, "2026-01-01T11:00:00Z")

	alert := alertApplication.AlertCacheDto{Fingerprint: fp1, Instance: instance1, Job: job1, StartsAt: startsAt1, AlertName: alertName1, Status: status1, Severity: severity1}
	return alert
}

func TestMatchRuleConditionUseCase_Valid(t *testing.T) {

	uc := application.NewMatchRuleConditionUseCase()

	alert := getAlert()
	condition := getCondition()

	res := uc.Execute(alert, condition)

	assert.True(t, res)
}

func TestMatchRuleConditionUseCase_FieldMissin(t *testing.T) {

	uc := application.NewMatchRuleConditionUseCase()

	alert := getAlert()
	condition := getCondition()
	condition.Field, _ = domain.NewConditionField("severity")
	condition.Value, _ = domain.NewConditionValue("critical")

	res := uc.Execute(alert, condition)

	assert.True(t, res)
}
