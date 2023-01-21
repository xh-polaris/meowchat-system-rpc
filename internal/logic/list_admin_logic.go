package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

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
	admins, count, err := l.svcCtx.AdminModel.ListAdmin(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var resp = make([]*pb.Admin, len(admins))
	for i, admin := range admins {
		resp[i] = svc.ConvertAdmin(admin)
	}

	return &pb.ListAdminResp{
		Admins: resp,
		Count:  count,
	}, nil
}
