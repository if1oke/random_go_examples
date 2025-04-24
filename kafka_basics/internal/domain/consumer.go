package domain

import "context"

type IConsumer interface {
	ReadMessage(ctx context.Context) (Message, error)
}
