package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn *amqp.Connection
}

func NewRabbit(connStr string) (*Rabbit, error) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}
	return &Rabbit{conn: conn}, nil
}

func (r *Rabbit) Channel() (*amqp.Channel, error) {
	return r.conn.Channel()
}

func (r *Rabbit) Close() error {
	return r.conn.Close()
}
