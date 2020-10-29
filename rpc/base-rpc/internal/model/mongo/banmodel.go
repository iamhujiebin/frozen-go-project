package mongo

import (
	"context"
	"frozen-go-project/rpc/base-rpc/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type (
	BanModel struct {
		globalConfig *config.Config
		client       *mongo.Client
		database     string
		collection   string
	}

	Ban struct {
		Id      primitive.ObjectID `bson:"_id,omitempty"`
		UserId  int64              `bson:"user_id,omitempty"`
		GuestId string             `bson:"guest_id,omitempty"`
	}
)

func NewCBanModel(client *mongo.Client, globalConfig *config.Config, database, collection string) *BanModel {
	return &BanModel{
		globalConfig: globalConfig,
		client:       client,
		database:     database,
		collection:   collection,
	}
}

func (m *BanModel) WithCollection(f func(col *mongo.Collection, ctx context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection(m.collection)
	f(collection, ctx)
}

func (m *BanModel) FindOne(userId int64, guestId string) (ban *Ban, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		where := bson.M{}
		if userId > 0 && len(guestId) > 0 {
			where = bson.M{
				"$or": []bson.M{
					{"user_id": userId},
					{"guest_id": guestId},
				},
			}
		}
		if userId > 0 && len(guestId) <= 0 {
			where = bson.M{
				"user_id": userId,
			}
		}
		if userId == 0 && len(guestId) > 0 {
			where = bson.M{
				"guest_id": guestId,
			}
		}
		if len(where) <= 0 {
			err = mongo.ErrNoDocuments
		}
		err = col.FindOne(ctx, where).Decode(&ban)
	})
	return
}
