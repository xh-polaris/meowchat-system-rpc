package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"

	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommunityLogic {
	return &DeleteCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommunityLogic) DeleteCommunity(in *pb.DeleteCommunityReq) (*pb.DeleteCommunityResp, error) {
	err := l.svcCtx.CommunityModel.DeleteCommunity(l.ctx, in.Id)

	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.DeleteCommunityResp{}, nil
}
