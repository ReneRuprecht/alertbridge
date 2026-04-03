package domain

import "context"

type AlertRepository interface {
	Save(context context.Context, alert Alert) error
	FindAlertsByInstance(context context.Context, instance string) ([]Alert, error)
}
