package application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/action/application"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockActionRepositoryFinder struct {
	actions domain.Action
	err     error
}

func (m *mockActionRepositoryFinder) FindByID(ctx context.Context, actionID domain.ActionID) (domain.Action, error) {
	return m.actions, m.err
}

func getSingleAction() domain.Action {
	id1, _ := domain.NewActionID()
	name1, _ := domain.NewActionName("webhook-url sender")
	desc1 := domain.NewActionDescription("sends message")
	cfg1, _ := domain.NewActionConfig(map[string]string{"url": "testurl"})
	t1, _ := domain.NewActionType("webhook")
	action := domain.Action{ID: id1, Name: name1, Description: desc1, Type: t1, Config: cfg1}

	return action
}

func TestListActionsById_Valid(t *testing.T) {

	data := getSingleAction()

	repo := &mockActionRepositoryFinder{
		actions: data,
		err:     nil,
	}

	uc := application.NewFindActionByIDUseCase(repo)

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	actionID := data.ID
	action, err := uc.Execute(ctx, actionID)

	require.NoError(t, err)

	assert.Equal(t, actionID, action.ID)
	assert.Equal(t, domain.ActionName("webhook-url sender"), action.Name)
	assert.Equal(t, domain.ActionDescription("sends message"), action.Description)
	assert.Equal(t, "testurl", action.Config["url"])
	assert.Equal(t, domain.ActionType("webhook"), action.Type)

}

func TestListActionsById_RepoError(t *testing.T) {

	repo := &mockActionRepositoryFinder{
		actions: domain.Action{},
		err:     errors.New("repo error"),
	}

	uc := application.NewFindActionByIDUseCase(repo)

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	actionID, _ := domain.NewActionID()
	_, err := uc.Execute(ctx, actionID)

	require.Error(t, err)
	assert.ErrorContains(t, err, "repo error")

}
