package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateNewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNewsLogic {
	return &CreateNewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateNewsLogic) CreateNews(in *pb.CreateNewsReq) (*pb.CreateNewsResp, error) {
	id := primitive.NewObjectID()

	err := l.svcCtx.NewsModel.Insert(l.ctx, &model.News{
		ID:          id,
		CommunityId: in.CommunityId,
		ImageUrl:    in.ImageUrl,
		LinkUrl:     in.LinkUrl,
		Type:        in.Type,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateNewsResp{
		Id: id.Hex(),
	}, nil
}
