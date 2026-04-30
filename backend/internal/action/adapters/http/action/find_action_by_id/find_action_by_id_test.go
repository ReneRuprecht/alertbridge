package findactionbyid

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
	"github.com/stretchr/testify/require"
)

const url = "/api/v1/actions"

type mockUseCase struct {
	action domain.Action
	err    error
}

func (m *mockUseCase) Execute(ctx context.Context, actionID domain.ActionID) (domain.Action, error) {
	return m.action, m.err
}

func getDemoAction(id domain.ActionID) domain.Action {

	name, _ := domain.NewActionName("slack sender")
	desc := domain.NewActionDescription("sends slack message")
	actionType, _ := domain.NewActionType("webhook")
	config, _ := domain.NewActionConfig(map[string]string{"url": "slack", "channel": "alert"})
	action := domain.Action{ID: id, Name: name, Description: desc, Type: actionType, Config: config}

	return action

}

func TestHandleFindActionByID_Valid(t *testing.T) {

	id, _ := domain.NewActionID()

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", uuid.UUID(id).String())

	w := httptest.NewRecorder()

	data := getDemoAction(id)

	uc := &mockUseCase{
		action: data,
		err:    nil,
	}
	handler := HandleFindActionByID(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var result findActionByIDResponse

	err := json.NewDecoder(w.Body).Decode(&result)
	require.NoError(t, err)

	assert.Equal(t, uuid.UUID(data.ID).String(), result.ID)
	assert.Equal(t, string(data.Name), result.Name)
	assert.Equal(t, string(data.Description), result.Description)
	assert.Equal(t, string(data.Type), result.Type)
	assert.Equal(t, "slack", result.Config["url"])
	assert.Equal(t, "alert", result.Config["channel"])

}

func TestHandleFindActionByID_InvalidID(t *testing.T) {

	id := "123"

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", id)

	w := httptest.NewRecorder()

	uc := &mockUseCase{
		action: domain.Action{},
		err:    nil,
	}
	handler := HandleFindActionByID(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}

}

func TestHandleFindActionByID_InternalError(t *testing.T) {

	id, _ := domain.NewActionID()

	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", uuid.UUID(id).String())

	w := httptest.NewRecorder()

	uc := &mockUseCase{
		action: domain.Action{},
		err:    errors.New("usecase error"),
	}
	handler := HandleFindActionByID(uc)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}

}
