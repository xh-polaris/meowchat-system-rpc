package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNewsLogic {
	return &DeleteNewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteNewsLogic) DeleteNews(in *pb.DeleteNewsReq) (*pb.DeleteNewsResp, error) {
	err := l.svcCtx.NewsModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.DeleteNewsResp{}, nil
}
