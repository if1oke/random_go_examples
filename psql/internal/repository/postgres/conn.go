package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Client struct {
	Pool *pgxpool.Pool
}

var DBUrl = "postgresql://postgres:12wqasxz@localhost:5432/sql_examls"

func NewClient() (*Client, error) {
	pool, err := pgxpool.Connect(context.Background(), DBUrl)
	if err != nil {
		return nil, errors.Errorf("could not connect to database: %v", err)
	}

	return &Client{Pool: pool}, nil
}
