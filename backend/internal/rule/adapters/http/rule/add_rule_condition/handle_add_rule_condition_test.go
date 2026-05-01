package addrulecondition

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/reneruprecht/alertbridge/backend/internal/rule/application"
	"github.com/reneruprecht/alertbridge/backend/internal/rule/domain"
)

const url = "/api/v1/rule/conditions"

type mockUseCase struct {
	err error
}

func (m *mockUseCase) Execute(ctx context.Context, condition domain.Condition) error {
	return m.err
}

func TestHandleAddRuleCondition_Created(t *testing.T) {
	body := `{
        "ruleID": "019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce",
        "name": "backend-critical-alerts",
        "operator": "equals",
        "field": "status",
        "value": "critical"
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := HandleAddRuleCondition(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}

func TestHandleAddRuleCondition_MapperError(t *testing.T) {
	body := `{
        "name": "backend-critical-alerts",
        "operator": "equals",
        "field": "status",
        "value": "critical"
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: nil,
	}
	handler := HandleAddRuleCondition(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleAddRuleCondition_RuleNotFound(t *testing.T) {
	body := `{
        "ruleID": "019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce",
        "name": "backend-critical-alerts",
        "operator": "equals",
        "field": "status",
        "value": "critical"
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: application.ErrorRuleNotFound,
	}
	handler := HandleAddRuleCondition(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestHandleAddRuleCondition_InternalError(t *testing.T) {
	body := `{
        "ruleID": "019de3d2-dfb4-7cf1-b56c-2b1e2b6f04ce",
        "name": "backend-critical-alerts",
        "operator": "equals",
        "field": "status",
        "value": "critical"
    }`

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		err: errors.New("repo error"),
	}
	handler := HandleAddRuleCondition(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}


