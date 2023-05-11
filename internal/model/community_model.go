package model

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CommunityCollectionName = "community"

var _ CommunityModel = (*CustomCommunityModel)(nil)

type (
	// CommunityModel is an interface to be customized, add more methods here,
	// and implement the added methods in CustomCommunityModel.
	CommunityModel interface {
		communityModel
		ListCommunity(ctx context.Context, req *pb.ListCommunityReq) ([]*Community, int64, error)
		DeleteCommunity(ctx context.Context, id string) error
	}

	CustomCommunityModel struct {
		*defaultCommunityModel
	}
)

func (c CustomCommunityModel) DeleteCommunity(ctx context.Context, id string) error {
	key := prefixCommunityCacheKey + id

	old := new(Community)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = c.conn.FindOneAndDelete(ctx, key, old, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	// delete children
	_, err = c.conn.DeleteMany(ctx, bson.M{
		"parentId": old.ID,
	})
	if err != nil {
		return err
	}

	// delete all cache
	err = c.conn.DelCache(ctx, prefixCommunityCacheKey+"*")
	if err != nil {
		return err
	}
	return nil
}

func (c CustomCommunityModel) ListCommunity(ctx context.Context, req *pb.ListCommunityReq) ([]*Community, int64, error) {
	var resp []*Community

	filter := bson.M{}
	if req.ParentId != "" {
		pid, err := primitive.ObjectIDFromHex(req.ParentId)
		if err == nil {
			filter["parentId"] = pid
		}
	}

	findOptions := ToFindOptions(req.Page, req.Size, req.Sort)

	err := c.conn.Find(ctx, &resp, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}

	count, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return resp, count, nil
}

// NewCommunityModel returns a model for the mongo.
func NewCommunityModel(url, db, collection string, c cache.CacheConf) CommunityModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &CustomCommunityModel{
		defaultCommunityModel: newDefaultCommunityModel(conn),
	}
}
