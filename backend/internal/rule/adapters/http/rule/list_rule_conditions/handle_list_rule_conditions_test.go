package listruleconditions

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

const url = "/api/v1/rule/conditions"

type mockUseCase struct {
	conditions []domain.Condition
	err        error
}

func (m *mockUseCase) Execute(ctx context.Context) ([]domain.Condition, error) {
	return m.conditions, m.err
}

func getCondition() domain.Condition {
	id, _ := domain.NewConditionID()
	name, _ := domain.NewConditionName("status critical")
	operator, _ := domain.NewConditionOperator("equals")
	field, _ := domain.NewConditionField("status")
	value, _ := domain.NewConditionValue("critical")

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

func TestHandleListRuleConditions_OK(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()

	condition := getCondition()
	uc := &mockUseCase{
		conditions: []domain.Condition{condition},
		err:        nil,
	}
	handler := HandleListRuleConditions(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var result ListRuleConditionsResponse

	json.NewDecoder(w.Body).Decode(&result)

	assert.Equal(t, 1, len(result.Conditions))

	conditionIDString := uuid.UUID(condition.ID).String()
	ruleIDString := uuid.UUID(condition.RuleID).String()
	assert.Equal(t, conditionIDString, result.Conditions[0].ID)
	assert.Equal(t, ruleIDString, result.Conditions[0].RuleID)
	assert.Equal(t, string(condition.Name), result.Conditions[0].Name)
	assert.Equal(t, string(condition.Operator), result.Conditions[0].Operator)
	assert.Equal(t, string(condition.Field), result.Conditions[0].Field)
	assert.Equal(t, string(condition.Value), result.Conditions[0].Value)
}

func TestHandleListRuleConditions_NotFound(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		conditions: []domain.Condition{},
		err:        application.ErrorConditionsNotFound,
	}
	handler := HandleListRuleConditions(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}

}

func TestHandleListRuleConditions_InternalError(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		conditions: []domain.Condition{},
		err:        errors.New("usecase error"),
	}
	handler := HandleListRuleConditions(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}

}
