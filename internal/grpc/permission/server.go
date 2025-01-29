package permission

import (
	"context"
	"errors"
	ssov1 "github.com/KonservnayaBaka/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sso/internal/services/permission"
)

const (
	emptyValue = 0
)

type Permission interface {
	AddPermission(ctx context.Context, user_id int, permission string) (bool, error)
	RemovePermission(ctx context.Context, user_id int, permission string) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedPermissionServer
	permission Permission
}

func Register(gRPC *grpc.Server, permission Permission) {
	ssov1.RegisterPermissionServer(gRPC, &serverAPI{permission: permission})
}

func (s *serverAPI) AddPermission(ctx context.Context, req *ssov1.AddPermissionRequest) (*ssov1.AddPermissionResponse, error) {
	err := validateAddPermission(req)
	if err != nil {
		return nil, err
	}

	success, err := s.permission.AddPermission(ctx, int(req.UserId), req.Permission)
	if err != nil {
		if errors.Is(err, permission.ErrPermissionExists) {
			return nil, status.Error(codes.AlreadyExists, "permission already exists")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.AddPermissionResponse{Success: success}, nil
}

func (s *serverAPI) RemovePermission(ctx context.Context, req *ssov1.RemovePermissionRequest) (*ssov1.RemovePermissionResponse, error) {
	err := validateRemovePermission(req)
	if err != nil {
		return nil, err
	}

	success, err := s.permission.RemovePermission(ctx, int(req.UserId), req.Permission)
	if err != nil {
		if errors.Is(err, permission.ErrPermissionDoesNotExist) {
			return nil, status.Error(codes.AlreadyExists, "permission doesnt exists")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.RemovePermissionResponse{Success: success}, nil
}
