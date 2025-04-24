package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"kafka-basics/internal/domain"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(addr, group, topic string) *Consumer {
	return &Consumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{addr},
			GroupID:  group,
			Topic:    topic,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
	}
}

func (c *Consumer) ReadMessage(ctx context.Context) (domain.Message, error) {
	m, err := c.Reader.ReadMessage(ctx)
	if err != nil {
		return domain.Message{}, err
	}

	var msg domain.Message
	err = json.Unmarshal(m.Value, &msg)
	if err != nil {
		return domain.Message{}, err
	}

	return msg, nil
}
