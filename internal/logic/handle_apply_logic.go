package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/common/constant"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

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

	apply, err := l.svcCtx.ApplyModel.FindOne(l.ctx, in.ApplyId)
	if err != nil {
		return nil, err
	}
	if in.IsRejected == true {
		apply.Status = constant.ApplyRejected
		apply.HandlerId = in.HandlerId
		_, err := l.svcCtx.ApplyModel.Update(l.ctx, apply)
		if err != nil {
			return nil, err
		}
	} else {
		apply.Status = constant.ApplyAccepted
		apply.HandlerId = in.HandlerId
		_, err := l.svcCtx.ApplyModel.Update(l.ctx, apply)
		if err != nil {
			return nil, err
		}
		userRole, err := l.svcCtx.UserRoleModel.FindOne(l.ctx, apply.ApplicantId)
		if err != nil {
			return nil, err
		}
		userRole.Roles = append(userRole.Roles, model.Role{
			Type:        constant.RoleCommunityAdmin,
			CommunityId: apply.CommunityId,
		})
		_, err = l.svcCtx.UserRoleModel.Update(l.ctx, userRole)
		if err != nil {
			return nil, err
		}
	}

	return &pb.HandleApplyResp{}, nil
}
