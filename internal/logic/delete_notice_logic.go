package logic

import (
	"context"
	"meowchat-notice-rpc/errorx"
	"meowchat-notice-rpc/internal/svc"
	"meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoticeLogic {
	return &DeleteNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteNoticeLogic) DeleteNotice(in *pb.DeleteNoticeReq) (*pb.DeleteNoticeResp, error) {
	err := l.svcCtx.NoticeModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.DeleteNoticeResp{}, nil
}
