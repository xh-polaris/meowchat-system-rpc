package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ApplyModel = (*customApplyModel)(nil)

type (
	// ApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplyModel.
	ApplyModel interface {
		applyModel
	}

	customApplyModel struct {
		*defaultApplyModel
	}
)

// NewApplyModel returns a model for the mongo.
func NewApplyModel(url, db, collection string) ApplyModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customApplyModel{
		defaultApplyModel: newDefaultApplyModel(conn),
	}
}
