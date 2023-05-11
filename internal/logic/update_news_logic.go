package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewsLogic {
	return &UpdateNewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNewsLogic) UpdateNews(in *pb.UpdateNewsReq) (*pb.UpdateNewsResp, error) {
	err := l.svcCtx.NewsModel.UpdateNews(l.ctx, in)

	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.UpdateNewsResp{}, nil
}
