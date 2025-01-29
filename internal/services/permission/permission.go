package permission

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sso/internal/lib/logger/sl"
	"strconv"
)

type Permission struct {
	log               *slog.Logger
	permissionControl PermissionControl
}

type PermissionControl interface {
	AddPermission(ctx context.Context, user_id int, permission string) (bool, error)
	RemovePermission(ctx context.Context, user_id int, permission string) (bool, error)
}

var (
	ErrPermissionExists       = errors.New("permission exists")
	ErrPermissionDoesNotExist = errors.New("ermission does not exist")
)

func New(
	log *slog.Logger,
	permissionControl PermissionControl,
) *Permission {
	return &Permission{
		log:               log,
		permissionControl: permissionControl,
	}
}

func (p *Permission) AddPermission(ctx context.Context, user_id int, permission string) (bool, error) {
	const op = "permission.AddPermission"

	log := p.log.With(
		slog.String("op", op),
		slog.String("user_id", strconv.Itoa(user_id)),
	)

	log.Info("permission in the process of adding")

	success, err := p.permissionControl.AddPermission(ctx, user_id, permission)
	if err != nil {
		if errors.Is(err, ErrPermissionExists) {
			log.Warn("permission already exists", sl.Err(err))

			return false, fmt.Errorf("%s: %w", op, ErrPermissionExists)
		}
		log.Error("failed tp added permission", sl.Err(err))

		return false, fmt.Errorf("%s, %w", op, err)
	}

	log.Info("permission added")

	return success, nil
}
func (p *Permission) RemovePermission(ctx context.Context, user_id int, permission string) (bool, error) {
	const op = "permission.RemovePermission"

	log := p.log.With(
		slog.String("op", op),
		slog.String("user_id", strconv.Itoa(user_id)),
	)

	success, err := p.permissionControl.RemovePermission(ctx, user_id, permission)
	if err != nil {
		if errors.Is(err, ErrPermissionDoesNotExist) {
			log.Warn("permission does not exist", sl.Err(err))

			return false, fmt.Errorf("%s: %w", op, ErrPermissionDoesNotExist)
		}
		log.Error("failed tp removed permission", sl.Err(err))

		return false, fmt.Errorf("%s, %w", op, err)
	}

	log.Info("permission removed")

	return success, nil
}
