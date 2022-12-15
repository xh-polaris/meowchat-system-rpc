package model

import (
	"context"
	"github.com/xh-polaris/meowchat-notice-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const AdminCollectionName = "admin"

var _ AdminModel = (*customAdminModel)(nil)

type (
	// AdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminModel.
	AdminModel interface {
		adminModel
		ListAdmin(ctx context.Context, query *pb.ListAdminReq) ([]*Admin, error)
		UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) error
	}

	customAdminModel struct {
		*defaultAdminModel
	}
)

func (m customAdminModel) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) error {
	key := prefixAdminCacheKey + req.Id

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return ErrInvalidObjectId
	}

	filter := bson.M{
		"_id": oid,
	}
	set := bson.M{
		"communityId": req.CommunityId,
		"name":        req.Name,
		"title":       req.Title,
		"phone":       req.Phone,
		"wechat":      req.Wechat,
		"avatarUrl":   req.AvatarUrl,
	}

	// 更新数据
	update := bson.M{
		"$set": set,
	}

	option := options.UpdateOptions{}
	option.SetUpsert(false)

	_, err = m.conn.UpdateOne(ctx, key, filter, update, &option)
	return err
}

func (m customAdminModel) ListAdmin(ctx context.Context, req *pb.ListAdminReq) ([]*Admin, error) {
	var resp []*Admin

	filter := bson.M{
		"communityId": req.CommunityId,
	}

	err := m.conn.Find(ctx, &resp, filter)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewAdminModel returns a newsmodel for the mongo.
func NewAdminModel(url, db, collection string, c cache.CacheConf) AdminModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customAdminModel{
		defaultAdminModel: newDefaultAdminModel(conn),
	}
}
