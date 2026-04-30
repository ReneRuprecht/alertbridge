package findactionbyid

type findActionByIDResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Config      map[string]string `json:"config"`
}
