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
	res, err := l.svcCtx.ApplyModel.FindAllApplyByCommunityId(l.ctx, in)
	if err != nil {
		return nil, err
	}
	apply := make([]*pb.Apply, 0, len(res))
	for _, x := range res {
		apply = append(apply, &pb.Apply{
			ApplyId:     x.ID.Hex(),
			ApplicantId: x.ApplicantId,
			CommunityId: x.CommunityId,
		})
	}
	return &pb.ListApplyResp{Apply: apply}, nil
}
