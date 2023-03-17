package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/common/constant"
	"github.com/xh-polaris/meowchat-system-rpc/common/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleApplyLogic {
	return &HandleApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandleApplyLogic) HandleApply(in *pb.HandleApplyReq) (*pb.HandleApplyResp, error) {
	if in.IsRejected {
		id, err := primitive.ObjectIDFromHex(in.ApplyId)
		if err != nil {
			return nil, errorx.ErrInvalidObjectId
		}
		roles, err := l.svcCtx.UserRoleModel.FindOne(l.ctx, in.ApplyId)
		if err != nil {
			return nil, err
		}
		apply, err := l.svcCtx.ApplyModel.FindOne(l.ctx, in.ApplyId)
		if err != nil {
			return nil, err
		}
		roles.Roles = append(roles.Roles, model.Role{
			Type:        constant.RoleNormalCommunityAdmin,
			CommunityId: apply.CommunityId,
		})
		_, err = l.svcCtx.UserRoleModel.Upsert(l.ctx, &model.UserRole{
			ID:    id,
			Roles: roles.Roles,
		})
		_, err = l.svcCtx.ApplyModel.Update(l.ctx, &model.Apply{
			IsRejected: 2,
			Handler:    in.MyId,
		})
		if err != nil {
			return nil, err
		}
	} else {
		_, err := l.svcCtx.ApplyModel.Update(l.ctx, &model.Apply{
			IsRejected: 3,
			Handler:    in.MyId,
		})
		if err != nil {
			return nil, err
		}
	}
	return &pb.HandleApplyResp{}, nil
}
