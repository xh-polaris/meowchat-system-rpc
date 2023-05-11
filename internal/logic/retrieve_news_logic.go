package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveNewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveNewsLogic {
	return &RetrieveNewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveNewsLogic) RetrieveNews(in *pb.RetrieveNewsReq) (*pb.RetrieveNewsResp, error) {
	news, err := l.svcCtx.NewsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.RetrieveNewsResp{
		News: svc.ConvertNews(news),
	}, nil
}
