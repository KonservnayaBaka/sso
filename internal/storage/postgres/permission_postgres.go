package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"sso/internal/storage"
)

func (s *Storage) AddPermission(ctx context.Context, user_id int, permission string) (bool, error) {
	const op = "storage.postgres.permission.AddPermission"

	_, err := s.pool.Exec(ctx, "INSERT INTO user_permissions(user_id, permission) VALUES($1, $2)", user_id, permission)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return false, fmt.Errorf("%s: %w", op, storage.ErrPermissionExists)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, nil
}

func (s *Storage) RemovePermission(ctx context.Context, user_id int, permission string) (bool, error) {
	const op = "storage.postgres.permission.RemovePermission"

	_, err := s.pool.Exec(ctx, "DELETE FROM user_permissions WHERE user_id = $1 AND permission = $2", user_id, permission)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, storage.ErrPermissionNotFound)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, nil
}
