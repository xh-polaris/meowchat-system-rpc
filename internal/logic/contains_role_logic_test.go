package logic

import (
	"context"
	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xh-polaris/meowchat-system-rpc/constant"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestContainsRoleLogic_ContainsRole(t *testing.T) {
	svcCtx := &svc.ServiceContext{
		UserRoleModel:  &model.CustomUserRoleModel{},
		CommunityModel: &model.CustomCommunityModel{},
	}
	l := NewContainsRoleLogic(context.Background(), svcCtx)

	PatchConvey("允许超级管理员", t, func() {
		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: []model.Role{
				{
					Type: constant.RoleSuperAdmin,
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: "123",
				},
				{
					Type: constant.RoleUser,
				},
			},
		}, nil).Build()
		defer mocker.UnPatch()

		result, err := l.ContainsRole(&pb.ContainsRoleReq{
			UserId:          "1",
			AllowSuperAdmin: true,
		})
		So(err, ShouldBeNil)
		So(result.Contains, ShouldBeTrue)
	})

	PatchConvey("不允许超级管理员1", t, func() {
		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: []model.Role{
				{
					Type: constant.RoleSuperAdmin,
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: "123",
				},
				{
					Type: constant.RoleUser,
				},
			},
		}, nil).Build()
		defer mocker.UnPatch()

		result, err := l.ContainsRole(&pb.ContainsRoleReq{
			UserId: "1",
			Role: &pb.Role{
				Type: "NotExist",
			},
			AllowSuperAdmin: false,
		})
		So(err, ShouldBeNil)
		So(result.Contains, ShouldBeFalse)
	})

	PatchConvey("只判断是否包含社区管理员，不判断社区ID", t, func() {
		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: []model.Role{
				{
					Type: constant.RoleSuperAdmin,
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: "any community id",
				},
				{
					Type: constant.RoleUser,
				},
			},
		}, nil).Build()
		defer mocker.UnPatch()

		result, err := l.ContainsRole(&pb.ContainsRoleReq{
			UserId: "1",
			Role: &pb.Role{
				Type: constant.RoleCommunityAdmin,
			},
		})
		So(err, ShouldBeNil)
		So(result.Contains, ShouldBeTrue)
	})

	PatchConvey("判断是否包含社区管理员并且判断社区ID，ID匹配", t, func() {
		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: []model.Role{
				{
					Type: constant.RoleSuperAdmin,
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: "63c8d68105b1466da4e62d72",
				},
				{
					Type: constant.RoleUser,
				},
			},
		}, nil).Build()
		defer mocker.UnPatch()

		result, err := l.ContainsRole(&pb.ContainsRoleReq{
			UserId: "1",
			Role: &pb.Role{
				Type:        constant.RoleCommunityAdmin,
				CommunityId: "63c8d68105b1466da4e62d72",
			},
		})
		So(err, ShouldBeNil)
		So(result.Contains, ShouldBeTrue)
	})

	PatchConvey("判断是否包含社区管理员并且判断社区ID，ID不匹配", t, func() {
		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: []model.Role{
				{
					Type: constant.RoleSuperAdmin,
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: "63c8d68105b1466da4e62d72",
				},
				{
					Type: constant.RoleUser,
				},
			},
		}, nil).Build()
		defer mocker.UnPatch()

		mocker2 := Mock((*model.CustomCommunityModel).FindOne).Return(nil, errorx.ErrNotFound).Build()
		defer mocker2.UnPatch()

		result, err := l.ContainsRole(&pb.ContainsRoleReq{
			UserId: "1",
			Role: &pb.Role{
				Type:        constant.RoleCommunityAdmin,
				CommunityId: "0000008105b1466da4e62d72",
			},
		})
		So(err, ShouldBeNil)
		So(result.Contains, ShouldBeFalse)
	})

	PatchConvey("判断是否包含社区管理员并且判断社区ID，允许子社区", t, func() {
		tmpId := primitive.NewObjectID()
		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: []model.Role{
				{
					Type: constant.RoleSuperAdmin,
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: tmpId.Hex(),
				},
				{
					Type: constant.RoleUser,
				},
			},
		}, nil).Build()
		defer mocker.UnPatch()

		mocker2 := Mock((*model.CustomCommunityModel).FindOne).Return(&model.Community{
			ParentId: tmpId,
		}, nil).Build()
		defer mocker2.UnPatch()

		result, err := l.ContainsRole(&pb.ContainsRoleReq{
			UserId: "1",
			Role: &pb.Role{
				Type:        constant.RoleCommunityAdmin,
				CommunityId: "0000008105b1466da4e62d72",
			},
		})
		So(err, ShouldBeNil)
		So(result.Contains, ShouldBeTrue)
	})

	PatchConvey("判断是否包含社区管理员并且判断社区ID，多社区", t, func() {
		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: []model.Role{
				{
					Type: constant.RoleSuperAdmin,
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: "63c8d68105b1466da4e62000",
				},
				{
					Type:        constant.RoleCommunityAdmin,
					CommunityId: "63c8d68105b1466da4e62d72",
				},
				{
					Type: constant.RoleUser,
				},
			},
		}, nil).Build()
		defer mocker.UnPatch()

		mocker2 := Mock((*model.CustomCommunityModel).FindOne).Return(&model.Community{
			ParentId: primitive.NilObjectID,
		}, nil).Build()
		defer mocker2.UnPatch()

		result, err := l.ContainsRole(&pb.ContainsRoleReq{
			UserId: "1",
			Role: &pb.Role{
				Type:        constant.RoleCommunityAdmin,
				CommunityId: "63c8d68105b1466da4e62d72",
			},
		})
		So(err, ShouldBeNil)
		So(result.Contains, ShouldBeTrue)

		result, _ = l.ContainsRole(&pb.ContainsRoleReq{
			UserId: "1",
			Role: &pb.Role{
				Type:        constant.RoleCommunityAdmin,
				CommunityId: "63c8d68105b1466da4e62000",
			},
		})
		So(result.Contains, ShouldBeTrue)
	})

	PatchConvey("判断允许其他角色", t, func() {
		roles := []model.Role{
			{
				Type: constant.RoleSuperAdmin,
			},
			{
				Type: constant.RoleCommunityAdmin,
			},
			{
				Type: constant.RoleUser,
			},
			{
				Type: "another type",
			},
		}

		mocker := Mock((*model.CustomUserRoleModel).FindOne).Return(&model.UserRole{
			Roles: roles,
		}, nil).Build()
		defer mocker.UnPatch()

		for _, role := range roles {
			result, err := l.ContainsRole(&pb.ContainsRoleReq{
				UserId: "1",
				Role: &pb.Role{
					Type: role.Type,
				},
			})
			So(err, ShouldBeNil)
			So(result.Contains, ShouldBeTrue)
		}
	})
}
