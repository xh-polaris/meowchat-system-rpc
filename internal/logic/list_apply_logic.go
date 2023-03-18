package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListApplyLogic {
	return &ListApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListApplyLogic) ListApply(in *pb.ListApplyReq) (*pb.ListApplyResp, error) {
	all, err := l.svcCtx.ApplyModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListApplyResp{ApplyId: all}, nil
}
