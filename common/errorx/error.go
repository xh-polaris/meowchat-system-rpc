package errorx

import (
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound                 = status.Error(10001, "no such element")
	ErrInvalidObjectId          = status.Error(10002, "invalid objectId")
	ErrCommunityIdNotFound      = status.Error(10003, "communityId not found")
	ErrChildCommunityNotAllowed = status.Error(10004, "child community not allowed")
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
