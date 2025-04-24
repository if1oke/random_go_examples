package usecase

import "kafka-basics/internal/domain"

type MessageProducerUseCase struct {
	producer domain.IProducer
}

func NewMessageProducerUseCase(producer domain.IProducer) *MessageProducerUseCase {
	return &MessageProducerUseCase{
		producer: producer,
	}
}

func (u *MessageProducerUseCase) SendMessage(m domain.Message) error {
	return u.producer.SendMessage(m)
}
