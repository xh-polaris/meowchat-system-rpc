package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/common/constant"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApplyLogic {
	return &CreateApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 申请成为管理员
func (l *CreateApplyLogic) CreateApply(in *pb.CreateApplyReq) (*pb.CreateApplyResp, error) {
	if err := l.svcCtx.ApplyModel.Insert(l.ctx, &model.Apply{
		UserId:  in.UserId,
		Status:  constant.UnProcessed,
		Handler: "",
	}); err != nil {
		return nil, err
	}

	return &pb.CreateApplyResp{}, nil
}
