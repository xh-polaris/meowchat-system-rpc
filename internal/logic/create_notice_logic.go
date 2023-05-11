package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNoticeLogic {
	return &CreateNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateNoticeLogic) CreateNotice(in *pb.CreateNoticeReq) (*pb.CreateNoticeResp, error) {
	id := primitive.NewObjectID()

	err := l.svcCtx.NoticeModel.Insert(l.ctx, &model.Notice{
		ID:          id,
		CommunityId: in.CommunityId,
		Text:        in.Text,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateNoticeResp{
		Id: id.Hex(),
	}, nil
}
