package usecase

import (
	"context"
	"errors"
	"psql/internal/domain"
	"psql/internal/service/metrics"
)

type RoomUseCase struct {
	repo domain.IRoomRepository
	tx   ITxManager
	m    metrics.IMetrics
}

func NewRoomUseCase(repo domain.IRoomRepository, tx ITxManager, m metrics.IMetrics) *RoomUseCase {
	return &RoomUseCase{repo: repo, tx: tx, m: m}
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
			u.m.IncCounter("reserveFailed")
			return err
		}
		if !isAvailable {
			u.m.IncCounter("reserveFailed")
			return errors.New("already reserved")
		}
		err = u.repo.SetReserve(ctx, id, true)
		if err != nil {
			u.m.IncCounter("reserveFailed")
			return err
		}

		u.m.IncCounter("reserveSuccess")
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
