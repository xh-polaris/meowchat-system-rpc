package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type News struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CommunityId string             `bson:"communityId,omitempty" json:"communityId,omitempty"`
	ImageUrl    string             `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"`
	LinkUrl     string             `bson:"linkUrl,omitempty" json:"linkUrl,omitempty"`
	Type        string             `bson:"type,omitempty" json:"type,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
