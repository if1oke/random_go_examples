package usecase

import (
	"context"
	"errors"
	"psql/internal/domain"
)

type RoomUseCase struct {
	repo domain.IRoomRepository
	tx   ITxManager
}

func NewRoomUseCase(repo domain.IRoomRepository, tx ITxManager) *RoomUseCase {
	return &RoomUseCase{repo: repo, tx: tx}
}

func (u *RoomUseCase) IsAvailable(ctx context.Context, id int) (bool, error) {
	isAvailable, err := u.repo.IsAvailableByID(ctx, id)
	if err != nil {
		return false, err
	}
	return isAvailable, nil
}

func (u *RoomUseCase) Reserve(ctx context.Context, id int) error {
	return u.tx.RunTx(ctx, func(ctx context.Context) error {
		isAvailable, err := u.repo.IsAvailableByID(ctx, id)
		if err != nil {
			return err
		}
		if !isAvailable {
			return errors.New("already reserved")
		}
		err = u.repo.SetReserve(ctx, id, true)
		if err != nil {
			return err
		}
		return nil
	})
}

func (u *RoomUseCase) UnsetReserve(ctx context.Context, id int) error {
	err := u.repo.SetReserve(ctx, id, false)
	if err != nil {
		return err
	}
	return nil
}
