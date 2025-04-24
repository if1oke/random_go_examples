package usecase

import (
	"context"
)

type ITxManager interface {
	RunTx(ctx context.Context, fn func(ctx context.Context) error) error
}
