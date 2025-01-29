package user_info

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sso/internal/domain/models"
	"sso/internal/lib/logger/sl"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserInfo struct {
	log      *slog.Logger
	userInfo UserInfoRepos
}

type UserInfoRepos interface {
	GetUserInfo(ctx context.Context, userId int) (models.User, error)
}

func New(log *slog.Logger, userInfo UserInfoRepos) *UserInfo {
	return &UserInfo{
		log:      log,
		userInfo: userInfo,
	}
}

func (u *UserInfo) GetUserInfo(ctx context.Context, userId int) (models.User, error) {
	const op = "services.user_info.GetUserInfo"

	log := u.log.With(
		slog.String("op", op),
		slog.Int("userId", userId),
	)

	log.Info("attempting to get info about user")

	user, err := u.userInfo.GetUserInfo(ctx, userId)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			log.Warn("user not found", sl.Err(err))
		}

		return models.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
	}

	log.Info("successfully got info about user")

	return user, nil
}
