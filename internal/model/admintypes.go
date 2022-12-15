package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CommunityId string             `bson:"communityId,omitempty" json:"communityId,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Phone       string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Wechat      string             `bson:"wechat,omitempty" json:"wechat,omitempty"`
	AvatarUrl   string             `bson:"avatarUrl,omitempty" json:"avatarUrl,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
