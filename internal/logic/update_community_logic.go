package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommunityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommunityLogic {
	return &UpdateCommunityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommunityLogic) UpdateCommunity(in *pb.UpdateCommunityReq) (*pb.UpdateCommunityResp, error) {
	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, errorx.ErrInvalidObjectId
	}

	parentId, err := l.svcCtx.CheckParentCommunityId(l.ctx, in.ParentId)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.CommunityModel.Update(l.ctx, &model.Community{
		ID:       id,
		Name:     in.Name,
		ParentId: parentId,
	})

	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.UpdateCommunityResp{}, nil
}
