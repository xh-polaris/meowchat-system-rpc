package logic

import (
	"context"
	"meowchat-notice-rpc/errorx"

	"meowchat-notice-rpc/internal/svc"
	"meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminLogic {
	return &UpdateAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAdminLogic) UpdateAdmin(in *pb.UpdateAdminReq) (*pb.UpdateAdminResp, error) {
	err := l.svcCtx.AdminModel.UpdateAdmin(l.ctx, in)

	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.UpdateAdminResp{}, nil
}
