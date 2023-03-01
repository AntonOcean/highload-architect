package amqp

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"kek/internal/config"
)

type Publisher struct {
	amqpDial   *amqp.Connection
	amqpDialCh *amqp.Channel
	config     config.RabbitMQ
}

func BuildPublisher(cfg *config.Config) (*Publisher, error) {
	amqpDial, err := amqp.Dial(cfg.RabbitMQ.DSN)
	if err != nil {
		return nil, err
	}

	ch, err := amqpDial.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"ws-only", // name
		"fanout",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		amqpDial:   amqpDial,
		amqpDialCh: ch,
		config:     cfg.RabbitMQ,
	}, nil
}

func (p *Publisher) Push(message *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         body,
	}

	if err := p.amqpDialCh.PublishWithContext(ctx, "", p.config.Queue, false, false, msg); err != nil {
		return err
	}

	if message.Key == string(PostEvent) {
		if err := p.amqpDialCh.PublishWithContext(ctx, "ws-only", "", false, false, msg); err != nil {
			return err
		}
	}

	return nil
}

func (p *Publisher) Close() error {
	if err := p.amqpDialCh.Close(); err != nil {
		return err
	}

	if err := p.amqpDial.Close(); err != nil {
		return err
	}

	return nil
}
