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

	if in.IsRejected == false {
		apply, err := l.svcCtx.ApplyModel.FindOne(l.ctx, in.ApplyId)
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
	}
	_, err := l.svcCtx.ApplyModel.Delete(l.ctx, in.ApplyId)
	if err != nil {
		return nil, err
	}
	return &pb.HandleApplyResp{}, nil
}
