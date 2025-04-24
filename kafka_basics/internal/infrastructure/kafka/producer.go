package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"kafka-basics/internal/domain"
	"log"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(addr string, topic string) *Producer {
	return &Producer{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(addr),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) SendMessage(m domain.Message) error {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = p.Writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: jsonBytes,
		},
	)
	if err != nil {
		log.Fatalf("error writing to kafka: %v\n", err)
	}

	fmt.Printf("Message sent to Kafka: %v\n", string(jsonBytes))
	return nil
}
