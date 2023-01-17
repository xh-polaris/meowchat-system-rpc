package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"

	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainsRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContainsRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainsRoleLogic {
	return &ContainsRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContainsRoleLogic) ContainsRole(in *pb.ContainsRoleReq) (*pb.ContainsRoleResp, error) {
	userRole, _ := l.svcCtx.UserRoleModel.FindOne(l.ctx, in.UserId)
	if userRole != nil {
		for _, role := range userRole.Roles {
			if role.Type == in.Role.Type {
				return &pb.ContainsRoleResp{
					Contains: in.Role.Type != model.RoleCommunityAdmin ||
						// 是CommunityAdmin 并且校验CommunityId的情况
						in.Role.CommunityId == "" ||
						role.CommunityId == in.Role.CommunityId,
				}, nil
			}
		}
	}

	return &pb.ContainsRoleResp{
		Contains: false,
	}, nil
}
