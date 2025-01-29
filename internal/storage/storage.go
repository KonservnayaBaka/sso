package storage

import "errors"

var (
	ErrUserExists       = errors.New("user already exists")
	ErrPermissionExists = errors.New("permission already exists")
	ErrUserNotFound     = errors.New("user not found")
	ErrAppNotFound      = errors.New("app not found")
)
