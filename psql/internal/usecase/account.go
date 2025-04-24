package usecase

import (
	"context"
	"errors"
	"psql/internal/domain"
)

type AccountUseCase struct {
	repo domain.IAccountRepository
	tx   ITxManager
}

func NewAccountUseCase(repo domain.IAccountRepository, tx ITxManager) *AccountUseCase {
	return &AccountUseCase{repo: repo, tx: tx}
}

func (u *AccountUseCase) Transfer(ctx context.Context, fromID, toID int, amount int) error {
	return u.tx.RunTx(ctx, func(ctx context.Context) error {
		from, err := u.repo.GetByID(ctx, fromID)
		if err != nil {
			return err
		}
		to, err := u.repo.GetByID(ctx, toID)
		if err != nil {
			return err
		}

		if from.Balance < amount {
			return errors.New("insufficient funds")
		}

		from.Balance -= amount
		to.Balance += amount

		if err = u.repo.UpdateBalance(ctx, from.ID, from.Balance); err != nil {
			return err
		}
		if err = u.repo.UpdateBalance(ctx, to.ID, to.Balance); err != nil {
			return err
		}

		return nil
	})
}
