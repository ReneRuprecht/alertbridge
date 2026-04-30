package application_test

import (
	"context"
	"testing"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/action/application"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockActionRepositoryReader struct {
	actions []domain.Action
	err     error
}

func (m *mockActionRepositoryReader) List(ctx context.Context) ([]domain.Action, error) {
	return m.actions, m.err
}

func getActions() []domain.Action {
	id1, _ := domain.NewActionID()
	name1, _ := domain.NewActionName("webhook-url sender")
	desc1 := domain.NewActionDescription("")
	cfg1, _ := domain.NewActionConfig(map[string]string{"url": "testurl"})
	t1, _ := domain.NewActionType("webhook")
	a1 := domain.Action{ID: id1, Name: name1, Description: desc1, Type: t1, Config: cfg1}

	id2, _ := domain.NewActionID()
	name2, _ := domain.NewActionName("slack sender")
	desc2 := domain.NewActionDescription("sends to slack channel")
	cfg2, _ := domain.NewActionConfig(map[string]string{"channel": "slack"})
	t2, _ := domain.NewActionType("webhook")
	a2 := domain.Action{ID: id2, Name: name2, Description: desc2, Type: t2, Config: cfg2}

	actions := []domain.Action{a1, a2}

	return actions
}

func TestListActionsUseCase_Valid(t *testing.T) {

	data := getActions()

	repo := &mockActionRepositoryReader{actions: data, err: nil}

	uc := application.NewListActionsUseCase(repo)

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	actions, err := uc.Execute(ctx)

	require.NoError(t, err)

	assert.Equal(t, 2, len(actions))
	assert.Equal(t, domain.ActionName("webhook-url sender"), actions[0].Name)
	assert.Equal(t, domain.ActionDescription(""), actions[0].Description)
	assert.Equal(t, domain.ActionType("webhook"), actions[0].Type)
	assert.Equal(t, "testurl", actions[0].Config["url"])
	assert.Equal(t, "slack", actions[1].Config["channel"])
	assert.Equal(t, domain.ActionName("slack sender"), actions[1].Name)
	assert.Equal(t, domain.ActionDescription("sends to slack channel"), actions[1].Description)

}

func TestListActionsUseCase_ValidEmpty(t *testing.T) {

	repo := &mockActionRepositoryReader{actions: []domain.Action{}}

	uc := application.NewListActionsUseCase(repo)

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	actions, err := uc.Execute(ctx)
	require.NoError(t, err)

	assert.Equal(t, 0, len(actions))

}
