package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/constant"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户的角色
func (l *UpdateUserRoleLogic) UpdateUserRole(in *pb.UpdateUserRoleReq) (*pb.UpdateUserRoleResp, error) {
	id, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, errorx.ErrInvalidObjectId
	}

	roles := make([]model.Role, len(in.Roles))
	for i, role := range in.Roles {
		if role.Type == constant.RoleCommunityAdmin {
			id, _ := l.svcCtx.CheckCommunityIdExist(l.ctx, role.CommunityId)
			if id == primitive.NilObjectID {
				return nil, errorx.ErrCommunityIdNotFound
			}
		}
		roles[i] = model.Role{
			Type:        role.Type,
			CommunityId: role.CommunityId,
		}
	}

	_, err = l.svcCtx.UserRoleModel.Upsert(l.ctx, &model.UserRole{
		ID:    id,
		Roles: roles,
	})
	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.UpdateUserRoleResp{}, nil
}
