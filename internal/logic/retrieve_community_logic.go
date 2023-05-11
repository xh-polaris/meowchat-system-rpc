package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveCommunityLogic {
	return &RetrieveCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveCommunityLogic) RetrieveCommunity(in *pb.RetrieveCommunityReq) (*pb.RetrieveCommunityResp, error) {
	community, err := l.svcCtx.CommunityModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.RetrieveCommunityResp{
		Community: svc.ConvertCommunity(community),
	}, nil
}
