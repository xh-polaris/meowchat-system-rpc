package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/constant"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainsRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContainsRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainsRoleLogic {
	return &ContainsRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 判断cid1的社区是不是cid2的社区的子社区
func (l *ContainsRoleLogic) subCommunityOf(cid1, cid2 string) bool {
	if cid1 == cid2 {
		return true
	}
	c1, _ := l.svcCtx.CommunityModel.FindOne(l.ctx, cid1)
	return c1 != nil && c1.ParentId.Hex() == cid2
}

func (l *ContainsRoleLogic) ContainsRole(in *pb.ContainsRoleReq) (resp *pb.ContainsRoleResp, err error) {
	resp = &pb.ContainsRoleResp{}

	if in.Role == nil {
		in.Role = &pb.Role{}
	}

	userRole, _ := l.svcCtx.UserRoleModel.FindOne(l.ctx, in.UserId)
	if userRole == nil {
		return
	}

	for _, role := range userRole.Roles {
		switch role.Type {
		case constant.RoleSuperAdmin:
			if in.AllowSuperAdmin || in.Role.Type == constant.RoleSuperAdmin {
				resp.Contains = true
				return
			}
		case constant.RoleCommunityAdmin:
			if in.Role.Type == constant.RoleCommunityAdmin &&
				(in.Role.CommunityId == "" || l.subCommunityOf(in.Role.CommunityId, role.CommunityId)) {
				resp.Contains = true
				return
			}
		default:
			if in.Role.Type == role.Type {
				resp.Contains = true
				return
			}
		}
	}

	return
}
