package createaction

type createActionRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Config      map[string]string `json:"config"`
}
