package alertmanager

type WebhookRequest struct {
	Alerts []AlertDto `json:"alerts"`
}
