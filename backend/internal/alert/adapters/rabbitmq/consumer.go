package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	application "github.com/reneruprecht/alertbridge/backend/internal/worker/application"
)

type Consumer struct {
	ch        *amqp.Channel
	queue     string
	processor application.AlertProcessorUseCase
}

func NewAlertEventConsumer(ch *amqp.Channel, queue string, processor application.AlertProcessorUseCase) *Consumer {
	return &Consumer{ch: ch, queue: queue, processor: processor}
}

func (c *Consumer) Consume(ctx context.Context) error {

	msgs, err := c.ch.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			alertID, err := toAlertFingerprint(string(msg.Body))

			if err != nil {
				log.Printf("mapping to fingerprint failed: %v", err)
				return
			}

			if err := c.processor.Execute(ctx, alertID); err != nil {
				log.Printf("process failed: %v", err)
			}
		}
	}()

	log.Printf("consumer started, waiting for messages")

	<-ctx.Done()
	return nil
}
