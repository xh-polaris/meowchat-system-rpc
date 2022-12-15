package errorx

import (
	"github.com/xh-polaris/meowchat-notice-rpc/internal/model"
	"google.golang.org/grpc/status"
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
