package rabbit

import (
	"context"
	"errors"

	"github.com/streadway/amqp"
)

type Messages struct {
	connectRabbitMQ *amqp.Connection
	channelRabbitMQ *amqp.Channel
	queues          map[string]struct{}
}

func New(amqpServerURL string, queues map[string]struct{}) (*Messages, error) {
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		return nil, err
	}
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return nil, err
	}

	for name, _ := range queues {
		_, err = channelRabbitMQ.QueueDeclare(
			name,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return nil, err
		}

	}

	return &Messages{connectRabbitMQ: connectRabbitMQ, channelRabbitMQ: channelRabbitMQ, queues: queues}, nil
}
func (m *Messages) Close(ctx context.Context) error {

	m.channelRabbitMQ.Close()
	return m.connectRabbitMQ.Close()
}
func (m *Messages) Put(queue string, data []byte) error {
	_, ok := m.queues[queue]

	if ok {
		if err := m.channelRabbitMQ.Publish(
			"",
			queue,
			false,
			false,
			amqp.Publishing{ContentType: "application/json", Body: data},
		); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("Key not found")
	}

}
func (m *Messages) Get(queue string) (<-chan amqp.Delivery, error) {
	_, ok := m.queues[queue]
	if ok {
		messages, err := m.channelRabbitMQ.Consume(
			queue,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return nil, err
		}
		return messages, nil
	} else {

		return nil, errors.New("Key not found")
	}
}
