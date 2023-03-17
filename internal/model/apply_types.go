package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apply struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId      string             `bson:"userId,omitempty" json:"userId,omitempty"`
	CommunityId string             `bson:"communityId,omitempty" json:"communityId,omitempty"`
	IsRejected  int                `bson:"isRejected,omitempty" json:"isrejected,omitempty"`
	Handler     string             `bson:"handler,omitempty" json:"handler,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
