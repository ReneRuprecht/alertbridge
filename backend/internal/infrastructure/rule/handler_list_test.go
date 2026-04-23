package rule

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	domain "github.com/reneruprecht/alertbridge/backend/internal/domain/rule"
	"github.com/stretchr/testify/assert"
)

const list_rule_url = "/api/v1/rule"

type mockListRulesUseCase struct {
	rules []domain.Rule
	err   error
}

func (m *mockListRulesUseCase) Execute(ctx context.Context) ([]domain.Rule, error) {
	return m.rules, m.err
}

func newListRulesHandler(rules []domain.Rule, error error) http.HandlerFunc {
	uc := &mockListRulesUseCase{
		err:   error,
		rules: rules,
	}
	handler := HandleListRules(uc)
	return handler
}

func TestHandleListRules_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()

	id := domain.NewRuleId()

	rules := []domain.Rule{domain.Rule{ID: id, Name: "backend-critical-alerts", Description: "testDesc", Priority: 100, Enabled: true}}

	handler := newListRulesHandler(rules, nil)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var result ListRulesResponse

	json.NewDecoder(w.Body).Decode(&result)

	assert.Equal(t, 1, len(result.Rules))

	assert.Equal(t, id.String(), result.Rules[0].ID)
	assert.Equal(t, "backend-critical-alerts", result.Rules[0].Name)
	assert.Equal(t, "testDesc", result.Rules[0].Description)
	assert.Equal(t, 100, result.Rules[0].Priority)
	assert.Equal(t, true, result.Rules[0].Enabled)

}

func TestHandleListRules_InternalServerError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()


	handler := newListRulesHandler(nil, errors.New("failed"))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}

}

