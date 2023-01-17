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

type CreateCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommunityLogic {
	return &CreateCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommunityLogic) CreateCommunity(in *pb.CreateCommunityReq) (*pb.CreateCommunityResp, error) {
	id := primitive.NewObjectID()

	parentId, err := l.svcCtx.CheckParentCommunityId(l.ctx, in.ParentId)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.CommunityModel.Insert(l.ctx, &model.Community{
		ID:       id,
		Name:     in.Name,
		ParentId: parentId,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateCommunityResp{
		Id: id.Hex(),
	}, nil
}
