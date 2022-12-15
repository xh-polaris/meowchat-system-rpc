package errorx

import (
	"google.golang.org/grpc/status"
	"meowchat-notice-rpc/internal/model"
)

var (
	ErrNotFound        = status.Error(10001, "no such element")
	ErrInvalidObjectId = status.Error(10002, "invalid objectId")
)

func Switch(err error) error {
	switch err {
	case model.ErrNotFound:
		return ErrNotFound
	case model.ErrInvalidObjectId:
		return ErrInvalidObjectId
	default:
		return err
	}
}
