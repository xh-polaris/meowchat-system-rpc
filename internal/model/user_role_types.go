package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Type        string `bson:"type,omitempty" json:"type,omitempty"`
	CommunityId string `bson:"communityId,omitempty" json:"communityId,omitempty"`
}

type UserRole struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Roles    []Role             `bson:"roles,omitempty" json:"roles,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
