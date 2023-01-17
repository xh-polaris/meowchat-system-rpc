package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"

	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNoticeLogic) UpdateNotice(in *pb.UpdateNoticeReq) (*pb.UpdateNoticeResp, error) {
	err := l.svcCtx.NoticeModel.UpdateNotice(l.ctx, in)

	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.UpdateNoticeResp{}, nil
}
