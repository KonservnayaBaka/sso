package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/config"
	"sso/internal/services/auth"
	"sso/internal/services/permission"
	"sso/internal/services/user_info"
	"sso/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, cfg *config.Config) *App {
	storage, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, cfg.TokenTTL)
	permissionService := permission.New(log, storage)
	userInfoSevice := user_info.New(log, storage)

	grpcApp := grpcapp.New(log, authService, permissionService, userInfoSevice, cfg.GRPC.Port)

	return &App{
		GRPCSrv: grpcApp,
	}
}
