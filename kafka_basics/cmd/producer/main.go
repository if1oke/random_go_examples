package main

import (
	"kafka-basics/internal/domain"
	"kafka-basics/internal/infrastructure/kafka"
	"kafka-basics/internal/usecase"
	"log"
)

func main() {
	producer := kafka.NewProducer("192.168.100.191:9092", "demo-test")
	defer producer.Writer.Close()

	useCase := usecase.NewMessageProducerUseCase(producer)

	message := domain.Message{
		ID:     1,
		Author: "Bob",
		Text:   "Hello World Kafka",
	}

	err := useCase.SendMessage(message)
	if err != nil {
		log.Fatalf("error sending message: %v", err)
	}
}
