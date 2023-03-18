package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ ApplyModel = (*customApplyModel)(nil)

const ApplyCollectionName = "apply"

type (
	// ApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplyModel.
	ApplyModel interface {
		applyModel
		FindAll(ctx context.Context) ([]string, error)
	}

	customApplyModel struct {
		*defaultApplyModel
	}
)

func (m customApplyModel) FindAll(ctx context.Context) ([]string, error) {
	var resp []*Apply
	err := m.conn.Find(ctx, &resp, bson.M{})
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(resp))
	for _, x := range resp {
		id := x.ID.Hex()
		res = append(res, id)
	}
	return res, nil
}

// NewApplyModel returns a model for the mongo.
func NewApplyModel(url, db, collection string) ApplyModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customApplyModel{
		defaultApplyModel: newDefaultApplyModel(conn),
	}
}
