package listconditionsbyruleid

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
	"gotest.tools/v3/assert"
)

const url = "/api/v1/rules/{id}/conditions"

type mockUseCase struct {
	response application.ListConditionsByRuleIDUseCaseDto
	err      error
}

func (m *mockUseCase) Execute(ctx context.Context, ruleID domain.RuleId) (application.ListConditionsByRuleIDUseCaseDto, error) {
	return m.response, m.err
}

func getCondition() domain.Condition {
	id, _ := domain.NewConditionID()
	name, _ := domain.NewConditionName("status critical")
	operator, _ := domain.NewConditionOperator("equals")
	field, _ := domain.NewConditionField("status")
	value, _ := domain.NewConditionValue("critical")

	ruleIDParsed, _ := uuid.Parse("019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce")
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

func getRule() domain.Rule {

	id, _ := uuid.Parse("019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce")
	name, _ := domain.NewRuleName("testName")
	description := "testDesc"
	priority, _ := domain.NewRulePriority(100)
	enabled := true

	rule := domain.Rule{
		ID:          domain.RuleId(id),
		Name:        name,
		Description: description,
		Priority:    priority,
		Enabled:     enabled,
	}
	return rule
}

func TestHandleListConditionsByRule_OK(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", "019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce")
	w := httptest.NewRecorder()

	condition := getCondition()
	rule := getRule()

	uc := &mockUseCase{
		response: application.ListConditionsByRuleIDUseCaseDto{
			Rule:       rule,
			Conditions: []domain.Condition{condition},
		},
		err: nil,
	}
	handler := HandleListConditionsByRuleID(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var result ListConditionsByRuleIDResponse

	json.NewDecoder(w.Body).Decode(&result)

	assert.Equal(t, 1, len(result.Conditions))

	conditionIDString := uuid.UUID(condition.ID).String()
	ruleIDString := uuid.UUID(condition.RuleID).String()
	assert.Equal(t, conditionIDString, result.Conditions[0].ID)

	assert.Equal(t, ruleIDString, result.Rule.ID)
	assert.Equal(t, string(condition.Name), result.Conditions[0].Name)
	assert.Equal(t, string(condition.Operator), result.Conditions[0].Operator)
	assert.Equal(t, string(condition.Field), result.Conditions[0].Field)
	assert.Equal(t, string(condition.Value), result.Conditions[0].Value)
}

func TestHandleListConditionsByRule_ConditionsNotFound(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", "019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce")
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		response: application.ListConditionsByRuleIDUseCaseDto{},
		err:      application.ErrorConditionsNotFound,
	}
	handler := HandleListConditionsByRuleID(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}

}

func TestHandleListConditionsByRule_RuleNotFound(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", "019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce")
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		response: application.ListConditionsByRuleIDUseCaseDto{},
		err:      application.ErrorRuleNotFound,
	}
	handler := HandleListConditionsByRuleID(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}

}

func TestHandleListConditionsByRule_InternalError(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", "019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce")
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		response: application.ListConditionsByRuleIDUseCaseDto{},
		err:      errors.New("usecase error"),
	}
	handler := HandleListConditionsByRuleID(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}

}

