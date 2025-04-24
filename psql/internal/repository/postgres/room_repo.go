package postgres

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	"psql/internal/domain"
)

type RoomRepo struct {
	db *pgxpool.Pool
}

func (a *RoomRepo) GetByID(ctx context.Context, id int) (*domain.Room, error) {
	row := a.db.QueryRow(ctx, "SELECT * FROM room WHERE id = $1", id)
	var room domain.Room
	if err := row.Scan(&room.ID, &room.Reserved, &room.Price); err != nil {
		return nil, err
	}
	return &room, nil
}

func (a *RoomRepo) SetReserve(ctx context.Context, id int, status bool) error {
	if tx, ok := ctx.Value(TxKey).(*sql.Tx); ok {
		_, err := tx.ExecContext(ctx, "UPDATE room SET reserved = $1 WHERE id = $2", status, id)
		return err
	}

	_, err := a.db.Exec(ctx, "UPDATE room SET reserved = $1 WHERE id = $2", status, id)
	return err
}

func (a *RoomRepo) IsAvailableByID(ctx context.Context, id int) (bool, error) {
	row := a.db.QueryRow(ctx, "SELECT reserved FROM room WHERE id = $1", id)
	var reserved bool
	if err := row.Scan(&reserved); err != nil {
		return false, err
	}
	return !reserved, nil
}

func NewRoomRepo(db *pgxpool.Pool) *RoomRepo {
	return &RoomRepo{db: db}
}
