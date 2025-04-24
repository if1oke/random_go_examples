package postgres

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	"psql/internal/domain"
)

type AccountRepo struct {
	db *pgxpool.Pool
}

func NewAccountRepo(db *pgxpool.Pool) *AccountRepo {
	return &AccountRepo{db: db}
}

func (a AccountRepo) GetByID(ctx context.Context, id int) (*domain.Account, error) {
	row := a.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id)
	var account domain.Account
	if err := row.Scan(&account.ID, &account.Balance); err != nil {
		return nil, err
	}

	return &account, nil
}

func (a AccountRepo) UpdateBalance(ctx context.Context, id int, newBalance int) error {
	if tx, ok := ctx.Value(TxKey).(*sql.Tx); ok {
		_, err := tx.ExecContext(ctx, "UPDATE users SET balance = $1 WHERE id = $2", newBalance, id)
		return err
	}

	_, err := a.db.Exec(ctx, "UPDATE users SET balance = $1 WHERE id = $2", newBalance, id)
	return err
}
