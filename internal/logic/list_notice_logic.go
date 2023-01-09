package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-notice-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNoticeLogic {
	return &ListNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListNoticeLogic) ListNotice(in *pb.ListNoticeReq) (*pb.ListNoticeResp, error) {
	notices, count, err := l.svcCtx.NoticeModel.ListNotice(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var resp = make([]*pb.Notice, len(notices))
	for i, n := range notices {
		resp[i] = &pb.Notice{
			Id:          n.ID.Hex(),
			CommunityId: n.CommunityId,
			Text:        n.Text,
			CreateAt:    n.CreateAt.UnixMilli(),
			UpdateAt:    n.UpdateAt.UnixMilli(),
		}
	}

	return &pb.ListNoticeResp{
		Notices: resp,
		Count:   count,
	}, nil
}
