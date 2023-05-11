package model

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/constant"
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
		FindMany(ctx context.Context, role string, communityId string) ([]*UserRole, error)
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

func (m CustomUserRoleModel) FindMany(ctx context.Context, role string, communityId string) ([]*UserRole, error) {
	var resp []*UserRole

	switch role {
	case constant.RoleSuperAdmin:
		err := m.conn.Find(ctx, &resp, bson.M{"roles.type": role})
		if err != nil {
			return nil, err
		}
		return resp, nil
	case constant.RoleCommunityAdmin:
		err := m.conn.Find(ctx, &resp, bson.M{"roles.type": role, "roles.communityId": communityId})
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		err := m.conn.Find(ctx, &resp, bson.M{"roles.type": role, "roles.communityId": communityId})
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

// NewUserRoleModel returns a model for the mongo.
func NewUserRoleModel(url, db, collection string, c cache.CacheConf) UserRoleModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &CustomUserRoleModel{
		defaultUserRoleModel: newDefaultUserRoleModel(conn),
	}
}
