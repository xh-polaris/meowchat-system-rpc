package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apply struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId   string             `bson:"userId,omitempty" json:"userId,omitempty""`
	Status   int8               `bson:"status,omitempty" json:"status,omitempty"`
	Handler  string             `bson:"handler,omitempty" json:"handler,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
