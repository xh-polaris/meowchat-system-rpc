// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var prefixNewsCacheKey = "cache:news:"

type newsModel interface {
	Insert(ctx context.Context, data *News) error
	FindOne(ctx context.Context, id string) (*News, error)
	Update(ctx context.Context, data *News) error
	Delete(ctx context.Context, id string) error
}

type defaultNewsModel struct {
	conn *monc.Model
}

func newDefaultNewsModel(conn *monc.Model) *defaultNewsModel {
	return &defaultNewsModel{conn: conn}
}

func (m *defaultNewsModel) Insert(ctx context.Context, data *News) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixNewsCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *defaultNewsModel) FindOne(ctx context.Context, id string) (*News, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data News
	key := prefixNewsCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNewsModel) Update(ctx context.Context, data *News) error {
	data.UpdateAt = time.Now()
	key := prefixNewsCacheKey + data.ID.Hex()
	_, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return err
}

func (m *defaultNewsModel) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}
	key := prefixNewsCacheKey + id
	_, err = m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
	return err
}
