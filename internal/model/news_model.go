package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"meowchat-notice-rpc/pb"
	"time"
)

const NewsCollectionName = "news"

var _ NewsModel = (*customNewsModel)(nil)

type (
	// NewsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsModel.
	NewsModel interface {
		newsModel
		UpdateNews(ctx context.Context, req *pb.UpdateNewsReq) error
		ListNews(ctx context.Context, req *pb.ListNewsReq) ([]*News, error)
	}

	customNewsModel struct {
		*defaultNewsModel
	}
)

func (m customNewsModel) ListNews(ctx context.Context, req *pb.ListNewsReq) ([]*News, error) {
	var resp []*News

	filter := bson.M{
		"communityId": req.CommunityId,
	}

	err := m.conn.Find(ctx, &resp, filter)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m customNewsModel) UpdateNews(ctx context.Context, req *pb.UpdateNewsReq) error {
	key := prefixNewsCacheKey + req.Id

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return ErrInvalidObjectId
	}

	filter := bson.M{
		"_id": oid,
	}
	set := bson.M{
		"type":     req.Type,
		"imageUrl": req.ImageUrl,
		"linkUrl":  req.LinkUrl,
		"updateAt": time.Now(),
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

// NewNewsModel returns a newsmodel for the mongo.
func NewNewsModel(url, db, collection string, c cache.CacheConf) NewsModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customNewsModel{
		defaultNewsModel: newDefaultNewsModel(conn),
	}
}
