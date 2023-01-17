package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommunityLogic {
	return &ListCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCommunityLogic) ListCommunity(in *pb.ListCommunityReq) (*pb.ListCommunityResp, error) {
	communities, count, err := l.svcCtx.CommunityModel.ListCommunity(l.ctx, in)
	if err != nil {
		return nil, errorx.Switch(err)
	}

	var resp = make([]*pb.Community, len(communities))
	for i, community := range communities {
		resp[i] = svc.ConvertCommunity(community)
	}

	return &pb.ListCommunityResp{
		Communities: resp,
		Count:       count,
	}, nil
}
