package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-notice-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAdminLogic {
	return &ListAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListAdminLogic) ListAdmin(in *pb.ListAdminReq) (*pb.ListAdminResp, error) {
	admins, err := l.svcCtx.AdminModel.ListAdmin(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var resp = make([]*pb.Admin, len(admins))
	for i, admin := range admins {
		resp[i] = &pb.Admin{
			Id:          admin.ID.Hex(),
			CommunityId: admin.CommunityId,
			Name:        admin.Name,
			Title:       admin.Title,
			Phone:       admin.Phone,
			Wechat:      admin.Wechat,
			AvatarUrl:   admin.AvatarUrl,
		}
	}

	return &pb.ListAdminResp{
		Admins: resp,
	}, nil
}
