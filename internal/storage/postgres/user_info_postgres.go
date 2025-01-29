package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"sso/internal/domain/models"
	"sso/internal/storage"
)

func (s *Storage) GetUserInfo(ctx context.Context, userId int) (models.User, error) {
	const op = "storage.postgres.permission.AddPermission"

	var user models.User

	query := `SELECT users.id, users.email, array_agg(user_permissions.permission)
			  FROM users
			  LEFT JOIN user_permissions
			  ON users.id=user_permissions.user_id
			  WHERE users.id = $1
			  GROUP BY users.id;`

	err := s.pool.QueryRow(ctx, query, userId).Scan(&user.ID, &user.Email, &user.UserPermissions)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
