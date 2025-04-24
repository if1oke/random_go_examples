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

func (u *MessageConsumerUseCase) Read(ctx context.Context) (domain.Message, error) {
	m, err := u.consumer.ReadMessage(ctx)
	if err != nil {
		return domain.Message{}, err
	}
	return m, nil
}
