package domain

import "context"

type IAccountRepository interface {
	GetByID(ctx context.Context, id int) (*Account, error)
	UpdateBalance(ctx context.Context, id int, newBalance int) error
}

type IRoomRepository interface {
	GetByID(ctx context.Context, id int) (*Room, error)
	SetReserve(ctx context.Context, id int, status bool) error
	IsAvailableByID(ctx context.Context, id int) (bool, error)
}
