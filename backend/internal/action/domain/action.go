package domain

type Action struct {
	ID          ActionID
	Name        ActionName
	Description ActionDescription
	Type        ActionType
	Config      ActionConfig
}
