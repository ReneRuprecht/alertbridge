package listaction

type listActionsResponse struct {
	Actions []action `json:"actions"`
}

type action struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Config      map[string]string `json:"config"`
}
