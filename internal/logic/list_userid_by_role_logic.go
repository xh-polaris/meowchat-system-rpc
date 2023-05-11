package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUseridByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUseridByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUseridByRoleLogic {
	return &ListUseridByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUseridByRoleLogic) ListUseridByRole(in *pb.ListUseridReq) (*pb.ListUseridResp, error) {

	Users, err := l.svcCtx.UserRoleModel.FindMany(l.ctx, in.RoleType, in.CommunityId)

	if err != nil {
		switch err {
		case model.ErrNotFound:
			return &pb.ListUseridResp{
				UserId: make([]string, 0),
			}, nil
		case model.ErrInvalidObjectId:
			return nil, errorx.ErrInvalidObjectId
		default:
			return nil, err
		}
	}

	var resp = make([]string, len(Users))
	for i, user := range Users {
		resp[i] = user.ID.Hex()
	}

	return &pb.ListUseridResp{
		UserId: resp,
	}, nil
}
