package domain

type IProducer interface {
	SendMessage(message Message) error
}
