package model

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

const ApplyCollectionName = "apply"

var _ ApplyModel = (*customApplyModel)(nil)

type (
	// ApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplyModel.
	ApplyModel interface {
		applyModel
		FindAllApplyByCommunityId(ctx context.Context, req *pb.ListApplyReq) ([]*Apply, error)
	}

	customApplyModel struct {
		*defaultApplyModel
	}
)

func (m customApplyModel) FindAllApplyByCommunityId(ctx context.Context, req *pb.ListApplyReq) ([]*Apply, error) {
	var resp []*Apply
	filter := bson.M{
		"communityId": req.CommunityId,
	}
	if err := m.conn.Find(ctx, &resp, filter); err != nil {
		return nil, err
	}
	return resp, nil
}

// NewApplyModel returns a model for the mongo.
func NewApplyModel(url, db, collection string) ApplyModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customApplyModel{
		defaultApplyModel: newDefaultApplyModel(conn),
	}
}
