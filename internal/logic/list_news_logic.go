package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-notice-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNewsLogic {
	return &ListNewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListNewsLogic) ListNews(in *pb.ListNewsReq) (*pb.ListNewsResp, error) {
	news, count, err := l.svcCtx.NewsModel.ListNews(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var resp = make([]*pb.News, len(news))
	for i, n := range news {
		resp[i] = &pb.News{
			Id:          n.ID.Hex(),
			CommunityId: n.CommunityId,
			ImageUrl:    n.ImageUrl,
			LinkUrl:     n.LinkUrl,
			Type:        n.Type,
		}
	}

	return &pb.ListNewsResp{
		News:  resp,
		Count: count,
	}, nil
}
