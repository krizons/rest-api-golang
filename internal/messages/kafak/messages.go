package kafak

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"
)

type Messages struct {
	brokerAddress string
	topic         string
	w             *kafka.Writer
	r             *kafka.Reader
}

func New(brokerAddress string, topic string) (*Messages, error) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	return &Messages{w: w, topic: topic, brokerAddress: brokerAddress}, nil
}
func (m *Messages) Close(ctx context.Context) error {
	<-ctx.Done()
	return m.w.Close()
}
func (m *Messages) Put(topic string, data []byte) error {
	if topic == m.topic {

		if err := m.w.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(topic),
			Value: data}); err != nil {
			return err
		}
	} else {
		return errors.New("Key not found")
	}
	return nil
}

func (m *Messages) Get(topic string, group string) ([]byte, error) {
	if m.r == nil {
		m.r = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{m.brokerAddress},
			Topic:   topic,
			GroupID: group,
		})
	}
	if topic == m.topic {
		msg, err := m.r.ReadMessage(context.Background())
		if err != nil {
			return nil, err
		}
		return msg.Value, nil
	} else {

		return nil, errors.New("Key not found")
	}

}
