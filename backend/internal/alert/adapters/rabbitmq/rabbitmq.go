package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type Publisher struct {
	ch    *amqp.Channel
	queue string
}

func NewAlertEventPublisher(ch *amqp.Channel, queue string) *Publisher {
	return &Publisher{ch: ch, queue: queue}
}

func (p *Publisher) Init() error {

	_, err := p.ch.QueueDeclare(
		p.queue,
		true,
		false,
		false,
		false,
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	return err
}

func (p *Publisher) Publish(ctx context.Context, alertID domain.Fingerprint) error {

	ctxPublish, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := alertID
	if err := p.ch.PublishWithContext(ctxPublish,
		"",
		p.queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		}); err != nil {
		return err
	}

	return nil

}
