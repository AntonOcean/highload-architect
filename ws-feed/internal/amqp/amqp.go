package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"

	"ws-feed/internal/domain"

	"ws-feed/internal/config"
	"ws-feed/internal/usecase"
)

type Consumer struct {
	amqpDial   *amqp.Connection
	amqpDialCh *amqp.Channel
	stop       chan bool
	config     config.RabbitMQ
}

func BuildConsumer(cfg *config.Config) (*Consumer, error) {
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

	return &Consumer{
		amqpDial:   amqpDial,
		stop:       make(chan bool),
		amqpDialCh: ch,
		config:     cfg.RabbitMQ,
	}, nil
}

func (c *Consumer) StartConsume(log *zap.Logger, uc usecase.ServiceUsecase) error {
	if err := c.amqpDialCh.Qos(1, 0, false); err != nil {
		return err
	}

	q, err := c.amqpDialCh.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = c.amqpDialCh.QueueBind(
		q.Name,    // queue name
		"",        // routing key
		"ws-only", // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	response, err := c.amqpDialCh.Consume(
		q.Name,
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

	wg := sync.WaitGroup{}

	for d := range response {
		var message Message

		if err := json.Unmarshal(d.Body, &message); err != nil {
			log.Error(fmt.Sprintf("err message format %+v", err))
			continue
		}

		ctx := context.Background()

		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			log.Info(fmt.Sprintf("start consume message %+v", msg))

			if KeyType(msg.Key) == PostEvent {
				data := domain.NewPost{}

				for k, v := range msg.Data {
					switch k {
					case "PostID":
						data.PostID = uuid.MustParse(v)
					case "AuthorID":
						data.AuthorID = uuid.MustParse(v)
					}
				}

				err = uc.UpdateFeed(ctx, data.PostID)
			}

			if err != nil {
				log.Error(fmt.Sprintf("err consume message %+v", err))
				return
			}
		}(message)
	}

	wg.Wait()
	c.stop <- true

	return nil
}

func (c *Consumer) Close(ctx context.Context) error {
	if err := c.amqpDialCh.Close(); err != nil {
		return err
	}

	if err := c.amqpDial.Close(); err != nil {
		return err
	}

	for {
		select {
		case <-c.stop:
			return nil
		case <-ctx.Done():
			return fmt.Errorf("some missing handlers :(")
		}
	}
}
