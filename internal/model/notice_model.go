package model

import (
	"context"
	"github.com/xh-polaris/meowchat-notice-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const NoticeCollectionName = "notice"

var _ NoticeModel = (*customNoticeModel)(nil)

type (
	// NoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticeModel.
	NoticeModel interface {
		noticeModel
		ListNotice(ctx context.Context, query *pb.ListNoticeReq) ([]*Notice, error)
		UpdateNotice(ctx context.Context, req *pb.UpdateNoticeReq) error
	}

	customNoticeModel struct {
		*defaultNoticeModel
	}
)

func (m customNoticeModel) UpdateNotice(ctx context.Context, req *pb.UpdateNoticeReq) error {
	key := prefixNoticeCacheKey + req.Id

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return ErrInvalidObjectId
	}

	filter := bson.M{
		"_id": oid,
	}

	update := bson.M{
		"$set": bson.M{
			"text": req.Text,
		},
	}
	_, err = m.conn.UpdateOne(ctx, key, filter, update)
	return err
}

func (m customNoticeModel) ListNotice(ctx context.Context, query *pb.ListNoticeReq) ([]*Notice, error) {
	var resp []*Notice

	filter := bson.M{
		"communityId": query.CommunityId,
	}

	err := m.conn.Find(ctx, &resp, filter)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewNoticeModel returns a noticemodel for the mongo.
func NewNoticeModel(url, db, collection string, c cache.CacheConf) NoticeModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customNoticeModel{
		defaultNoticeModel: newDefaultNoticeModel(conn),
	}
}
