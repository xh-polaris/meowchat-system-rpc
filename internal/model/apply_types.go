package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apply struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ApplicantId string             `bson:"applicantId,omitempty" json:"applicantId,omitempty"`
	CommunityId string             `bson:"communityId,omitempty" json:"communityId,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
