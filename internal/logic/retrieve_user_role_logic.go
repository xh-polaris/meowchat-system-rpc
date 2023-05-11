package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveUserRoleLogic {
	return &RetrieveUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户的所有角色
func (l *RetrieveUserRoleLogic) RetrieveUserRole(in *pb.RetrieveUserRoleReq) (*pb.RetrieveUserRoleResp, error) {
	userRole, err := l.svcCtx.UserRoleModel.FindOne(l.ctx, in.UserId)

	if err != nil {
		switch err {
		case model.ErrNotFound:
			return &pb.RetrieveUserRoleResp{
				Roles: make([]*pb.Role, 0),
			}, nil
		case model.ErrInvalidObjectId:
			return nil, errorx.ErrInvalidObjectId
		default:
			return nil, err
		}
	}

	var resp = make([]*pb.Role, len(userRole.Roles))
	for i, role := range userRole.Roles {
		resp[i] = &pb.Role{
			Type:        role.Type,
			CommunityId: role.CommunityId,
		}
	}

	return &pb.RetrieveUserRoleResp{
		Roles: resp,
	}, nil
}
