package domain

type AlertRepository interface {
	Save(alert Alert) error
}
