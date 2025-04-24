package main

import (
	"fmt"
	"kafka-basics/internal/infrastructure/kafka"
	"kafka-basics/internal/usecase"
	"log"
)

func main() {
	consumer := kafka.NewConsumer(
		"192.168.100.191:9092",
		"base_g1",
		"demo-test",
	)
	defer consumer.Reader.Close()

	useCase := usecase.NewMessageConsumerUseCase(consumer)

	for {
		m, err := useCase.Read()
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}
		fmt.Printf("📥 Received:  %+v\n", m)
	}
}
