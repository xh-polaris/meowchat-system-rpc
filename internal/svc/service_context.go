package svc

import (
	"github.com/xh-polaris/meowchat-system-rpc/internal/config"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	model.NewsModel
	model.AdminModel
	model.NoticeModel
	model.CommunityModel
	model.UserRoleModel
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		NewsModel:      model.NewNewsModel(c.Mongo.URL, c.Mongo.DB, model.NewsCollectionName, c.CacheConf),
		AdminModel:     model.NewAdminModel(c.Mongo.URL, c.Mongo.DB, model.AdminCollectionName, c.CacheConf),
		NoticeModel:    model.NewNoticeModel(c.Mongo.URL, c.Mongo.DB, model.NoticeCollectionName, c.CacheConf),
		CommunityModel: model.NewCommunityModel(c.Mongo.URL, c.Mongo.DB, model.CommunityCollectionName, c.CacheConf),
		UserRoleModel:  model.NewUserRoleModel(c.Mongo.URL, c.Mongo.DB, model.UserRoleCollectionName, c.CacheConf),
		Redis:          c.Redis.NewRedis(),
	}
}
