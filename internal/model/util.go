package model

import (
	"github.com/zeromicro/go-zero/core/mathx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ToFindOptions(page, size int64, sort map[string]int32) *options.FindOptions {
	page = int64(mathx.MaxInt(1, int(page)))
	findOptions := new(options.FindOptions)
	if size > 0 {
		findOptions.SetLimit(size)
		findOptions.SetSkip((page - 1) * size)
	}
	sortMap := sort
	if len(sortMap) > 0 {
		sort := bson.D{}
		for k, v := range sortMap {
			sort = append(sort, bson.E{Key: k, Value: v})
		}
		findOptions.SetSort(sort)
	}
	return findOptions
}
