package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const UserRoleCollectionName = "user_role"

var _ UserRoleModel = (*CustomUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in CustomUserRoleModel.
	UserRoleModel interface {
		userRoleModel
		Upsert(ctx context.Context, data *UserRole) (*mongo.UpdateResult, error)
	}

	CustomUserRoleModel struct {
		*defaultUserRoleModel
	}
)

func (m CustomUserRoleModel) Upsert(ctx context.Context, data *UserRole) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()
	key := prefixUserRoleCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(ctx, key,
		bson.M{"_id": data.ID},
		bson.M{"$set": data, "$setOnInsert": bson.M{"createAt": time.Now()}},
		&options.UpdateOptions{
			Upsert: &[]bool{true}[0],
		})
	return res, err
}

// NewUserRoleModel returns a model for the mongo.
func NewUserRoleModel(url, db, collection string, c cache.CacheConf) UserRoleModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &CustomUserRoleModel{
		defaultUserRoleModel: newDefaultUserRoleModel(conn),
	}
}
