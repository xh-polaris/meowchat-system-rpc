package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-notice-rpc/errorx"
	"github.com/xh-polaris/meowchat-notice-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveNoticeLogic {
	return &RetrieveNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveNoticeLogic) RetrieveNotice(in *pb.RetrieveNoticeReq) (*pb.RetrieveNoticeResp, error) {
	notice, err := l.svcCtx.NoticeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.RetrieveNoticeResp{
		Notice: &pb.Notice{
			Id:          notice.ID.Hex(),
			CommunityId: notice.CommunityId,
			Text:        notice.Text,
			CreateAt:    notice.CreateAt.UnixMilli(),
			UpdateAt:    notice.UpdateAt.UnixMilli(),
		},
	}, nil
}
