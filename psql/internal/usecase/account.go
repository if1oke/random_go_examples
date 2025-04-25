package usecase

import (
	"context"
	"errors"
	"psql/internal/domain"
	"psql/internal/service/metrics"
)

type AccountUseCase struct {
	repo domain.IAccountRepository
	tx   ITxManager
	m    metrics.IMetrics
}

func NewAccountUseCase(repo domain.IAccountRepository, tx ITxManager, m metrics.IMetrics) *AccountUseCase {
	return &AccountUseCase{repo: repo, tx: tx, m: m}
}

func (u *AccountUseCase) Transfer(ctx context.Context, fromID, toID int, amount int) error {
	return u.tx.RunTx(ctx, func(ctx context.Context) error {
		from, err := u.repo.GetByID(ctx, fromID)
		if err != nil {
			u.m.IncCounter("transferFailed")
			return err
		}
		to, err := u.repo.GetByID(ctx, toID)
		if err != nil {
			u.m.IncCounter("transferFailed")
			return err
		}

		if from.Balance < amount {
			u.m.IncCounter("transferFailed")
			return errors.New("insufficient funds")
		}

		from.Balance -= amount
		to.Balance += amount

		if err = u.repo.UpdateBalance(ctx, from.ID, from.Balance); err != nil {
			u.m.IncCounter("transferFailed")
			return err
		}
		if err = u.repo.UpdateBalance(ctx, to.ID, to.Balance); err != nil {
			u.m.IncCounter("transferFailed")
			return err
		}

		u.m.IncCounter("transferSuccess")

		return nil
	})
}
