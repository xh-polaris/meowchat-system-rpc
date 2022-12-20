package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CommunityId string             `bson:"communityId,omitempty" json:"communityId,omitempty"`
	Text        string             `bson:"text,omitempty" json:"text,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
