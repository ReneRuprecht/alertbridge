package listaction

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
)

const url = "/api/v1/actions"

type mockUseCase struct {
	actions []domain.Action
	err     error
}

func (m *mockUseCase) Execute(ctx context.Context) ([]domain.Action, error) {
	return m.actions, m.err
}

func getDemoActions() []domain.Action {

	id1, _ := domain.NewActionID()
	name1, _ := domain.NewActionName("slack sender")
	desc1 := domain.NewActionDescription("sends slack message")
	actionType1, _ := domain.NewActionType("webhook")
	config1, _ := domain.NewActionConfig(map[string]string{"url": "slack", "channel": "alert"})
	action1 := domain.Action{ID: id1, Name: name1, Description: desc1, Type: actionType1, Config: config1}

	id2, _ := domain.NewActionID()
	name2, _ := domain.NewActionName("email sender")
	desc2 := domain.NewActionDescription("sends email message")
	actionType2, _ := domain.NewActionType("webhook")
	config2, _ := domain.NewActionConfig(map[string]string{"to": "email", "signatur": "alert"})
	action2 := domain.Action{ID: id2, Name: name2, Description: desc2, Type: actionType2, Config: config2}

	return []domain.Action{action1, action2}

}

func TestHandleListActions_Valid(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, url, nil)
	w := httptest.NewRecorder()

	data := getDemoActions()

	uc := &mockUseCase{
		actions: data,
		err:     nil,
	}
	handler := HandleListActions(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var result listActionsResponse

	json.NewDecoder(w.Body).Decode(&result)

	assert.Equal(t, 2, len(result.Actions))

	assert.Equal(t, uuid.UUID(data[0].ID).String(), result.Actions[0].ID)
	assert.Equal(t, string(data[0].Name), result.Actions[0].Name)
	assert.Equal(t, string(data[0].Description), result.Actions[0].Description)
	assert.Equal(t, string(data[0].Type), result.Actions[0].Type)
	assert.Equal(t, "slack", result.Actions[0].Config["url"])
	assert.Equal(t, "alert", result.Actions[0].Config["channel"])
}

func TestHandleListActions_InternalError(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, url, nil)
	w := httptest.NewRecorder()

	uc := &mockUseCase{
		actions: []domain.Action{},
		err:     errors.New("usecase error"),
	}
	handler := HandleListActions(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}

}
