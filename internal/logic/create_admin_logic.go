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

type CreateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAdminLogic {
	return &CreateAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAdminLogic) CreateAdmin(in *pb.CreateAdminReq) (*pb.CreateAdminResp, error) {
	id := primitive.NewObjectID()

	err := l.svcCtx.AdminModel.Insert(l.ctx, &model.Admin{
		ID:          id,
		CommunityId: in.CommunityId,
		Name:        in.Name,
		Title:       in.Title,
		Phone:       in.Phone,
		Wechat:      in.Wechat,
		AvatarUrl:   in.AvatarUrl,
		UpdateAt:    time.Now(),
		CreateAt:    time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateAdminResp{
		Id: id.Hex(),
	}, nil
}
