package svc

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertAdmin(in *model.Admin) *pb.Admin {
	return &pb.Admin{
		Id:          in.ID.Hex(),
		CommunityId: in.CommunityId,
		Name:        in.Name,
		Title:       in.Title,
		Phone:       in.Phone,
		Wechat:      in.Wechat,
		AvatarUrl:   in.AvatarUrl,
	}
}

func ConvertNotice(in *model.Notice) *pb.Notice {
	return &pb.Notice{
		Id:          in.ID.Hex(),
		CommunityId: in.CommunityId,
		Text:        in.Text,
		CreateAt:    in.CreateAt.Unix(),
		UpdateAt:    in.UpdateAt.Unix(),
	}
}

func ConvertNews(in *model.News) *pb.News {
	return &pb.News{
		Id:          in.ID.Hex(),
		CommunityId: in.CommunityId,
		ImageUrl:    in.ImageUrl,
		LinkUrl:     in.LinkUrl,
		Type:        in.Type,
	}
}

func ConvertCommunity(in *model.Community) *pb.Community {
	pid := ""
	if in.ParentId != primitive.NilObjectID {
		pid = in.ParentId.Hex()
	}

	return &pb.Community{
		Id:       in.ID.Hex(),
		Name:     in.Name,
		ParentId: pid,
	}
}

func (s *ServiceContext) CheckCommunityIdExist(ctx context.Context, id string) (primitive.ObjectID, error) {
	if id == "" {
		return primitive.NilObjectID, nil
	}
	r, err := s.CommunityModel.FindOne(ctx, id)
	if err != nil {
		return primitive.NilObjectID, errorx.ErrCommunityIdNotFound
	}
	return r.ID, nil
}

func (s *ServiceContext) CheckParentCommunityId(ctx context.Context, parentId string) (primitive.ObjectID, error) {
	if parentId == "" {
		return primitive.NilObjectID, nil
	}
	r, err := s.CommunityModel.FindOne(ctx, parentId)
	if err != nil {
		return primitive.NilObjectID, errorx.ErrCommunityIdNotFound
	}
	if r.ParentId != primitive.NilObjectID {
		return primitive.NilObjectID, errorx.ErrChildCommunityNotAllowed
	}
	return r.ID, nil
}
