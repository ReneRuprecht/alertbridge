package rule

type ListRulesResponse struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Enabled     bool   `json:"enabled"`
}
