package usecase

import (
	"context"
	"kafka-basics/internal/domain"
)

type MessageConsumerUseCase struct {
	consumer domain.IConsumer
}

func NewMessageConsumerUseCase(consumer domain.IConsumer) *MessageConsumerUseCase {
	return &MessageConsumerUseCase{
		consumer: consumer,
	}
}

func (u *MessageConsumerUseCase) Read() (domain.Message, error) {
	m, err := u.consumer.ReadMessage(context.Background())
	if err != nil {
		return domain.Message{}, err
	}
	return m, nil
}
