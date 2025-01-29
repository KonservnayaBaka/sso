package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"sso/internal/config"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(cfg *config.Config) (*Storage, error) {
	const op = "storage.postgres.New"

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User, os.Getenv("DB_PASSWORD"),
		cfg.Database.Host, cfg.Database.Port,
		cfg.Database.DBName, cfg.Database.SslMode,
	)

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{pool: db}, nil
}
