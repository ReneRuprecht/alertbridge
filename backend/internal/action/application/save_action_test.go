package application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/reneruprecht/alertbridge/backend/internal/action/application"
	"github.com/reneruprecht/alertbridge/backend/internal/action/domain"
	"github.com/stretchr/testify/require"
)

type mockActionRepositoryWriter struct {
	err error
}

func (m *mockActionRepositoryWriter) Save(ctx context.Context, action domain.Action) error {
	return m.err
}

func getActionToSave() domain.Action {
	id1, _ := domain.NewActionID()
	cfg1, _ := domain.NewActionConfig(map[string]string{"url": "testurl"})
	t1, _ := domain.NewActionType("webhook")
	action := domain.Action{ID: id1, Type: t1, Config: cfg1}

	return action
}

func TestSaveActionUseCase_NoError(t *testing.T) {

	repo := &mockActionRepositoryWriter{err: nil}

	uc := application.NewSaveActionUseCase(repo)
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	action := getActionToSave()
	err := uc.Execute(ctx, action)

	require.NoError(t, err)
}

func TestSaveActionUseCase_RepoError(t *testing.T) {

	repo := &mockActionRepositoryWriter{err: errors.New("repo error")}

	uc := application.NewSaveActionUseCase(repo)
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	action := getActionToSave()
	err := uc.Execute(ctx, action)

	require.ErrorContains(t, err, "repo error")
}
