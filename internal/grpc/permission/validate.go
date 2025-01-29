package permission

import (
	ssov1 "github.com/KonservnayaBaka/protos/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateAddPermission(req *ssov1.AddPermissionRequest) error {
	if req.UserId == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}

	if req.Permission == "" {
		return status.Error(codes.InvalidArgument, "permission is required")
	}
	return nil
}

func validateRemovePermission(req *ssov1.RemovePermissionRequest) error {
	if req.UserId == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}

	if req.Permission == "" {
		return status.Error(codes.InvalidArgument, "permission is required")
	}
	return nil
}
