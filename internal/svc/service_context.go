package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"meowchat-notice-rpc/internal/config"
	"meowchat-notice-rpc/internal/model"
)

type ServiceContext struct {
	Config config.Config
	model.NewsModel
	model.AdminModel
	model.NoticeModel
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		NewsModel:   model.NewNewsModel(c.Mongo.URL, c.Mongo.DB, model.NewsCollectionName, c.CacheConf),
		AdminModel:  model.NewAdminModel(c.Mongo.URL, c.Mongo.DB, model.AdminCollectionName, c.CacheConf),
		NoticeModel: model.NewNoticeModel(c.Mongo.URL, c.Mongo.DB, model.NoticeCollectionName, c.CacheConf),
		Redis:       c.Redis.NewRedis(),
	}
}
