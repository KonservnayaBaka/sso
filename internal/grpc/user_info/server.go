package user_info

import (
	"context"
	"errors"
	ssov1 "github.com/KonservnayaBaka/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sso/internal/domain/models"
	"sso/internal/services/user_info"
)

type UserInfo interface {
	GetUserInfo(ctx context.Context, userId int) (models.User, error)
}

type serverAPI struct {
	ssov1.UnimplementedUserInfoServer
	userInfo UserInfo
}

func Register(gRPCServer *grpc.Server, userInfo UserInfo) {
	ssov1.RegisterUserInfoServer(gRPCServer, &serverAPI{userInfo: userInfo})
}

const emptyValue = 0

func (s *serverAPI) GetUserInfo(ctx context.Context, req *ssov1.GetUserInfoRequest) (*ssov1.GetUserInfoResponse, error) {
	if req.UserId == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	user, err := s.userInfo.GetUserInfo(ctx, int(req.UserId))
	if err != nil {
		if errors.Is(err, user_info.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &ssov1.GetUserInfoResponse{UserId: user.ID, Email: user.Email, Permissions: user.UserPermissions}, nil
}
