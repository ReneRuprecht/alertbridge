package action

import (
	"net/http"

	httpActionHandleCreateAction "github.com/reneruprecht/alertbridge/backend/internal/action/adapters/http/action/create_action"
	httpActionHandleFindActionByID "github.com/reneruprecht/alertbridge/backend/internal/action/adapters/http/action/find_action_by_id"
	httpActionHandleListActions "github.com/reneruprecht/alertbridge/backend/internal/action/adapters/http/action/list_actions"
	"github.com/reneruprecht/alertbridge/backend/internal/action/adapters/postgres"
	"github.com/reneruprecht/alertbridge/backend/internal/action/application"
	"github.com/reneruprecht/alertbridge/backend/internal/platform/postgres_db"
)

type ActionModule struct {
	SaveAction     application.SaveActionUseCase
	ListActions    application.ListActionsUseCase
	FindActionByID application.FindActionByIDUseCase
}

func NewActionModule(queries *postgres_db.Queries) *ActionModule {
	repo := postgres.NewActionRepository(queries)

	return &ActionModule{
		SaveAction:     application.NewSaveActionUseCase(repo),
		ListActions:    application.NewListActionsUseCase(repo),
		FindActionByID: application.NewFindActionByIDUseCase(repo),
	}
}

func (m *ActionModule) RegisterAlertRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /api/v1/actions", httpActionHandleCreateAction.HandleCreateAction(m.SaveAction))

	mux.HandleFunc("GET /api/v1/actions", httpActionHandleListActions.HandleListActions(m.ListActions))
	mux.HandleFunc("GET /api/v1/actions/{id}", httpActionHandleFindActionByID.HandleFindActionByID(m.FindActionByID))

}
