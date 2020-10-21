package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Ids struct {
		ObjectId  primitive.ObjectID `bson:"_id,omitempty"`
		TableName string             `bson:"tablename"`
		Id        int64              `bson:"id"`
	}
)
