package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveAdminLogic {
	return &RetrieveAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveAdminLogic) RetrieveAdmin(in *pb.RetrieveAdminReq) (*pb.RetrieveAdminResp, error) {
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.RetrieveAdminResp{
		Admin: svc.ConvertAdmin(admin),
	}, nil
}
