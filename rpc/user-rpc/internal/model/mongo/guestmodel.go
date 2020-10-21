package mongo

import (
	"context"
	"frozen-go-project/common/errors/db_errors"
	"frozen-go-project/rpc/user-rpc/internal/config"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type (
	GuestsModel struct {
		globalConfig *config.Config
		client       *mongo.Client
		database     string
		collection   string
	}

	Guests struct {
		Id          primitive.ObjectID `bson:"_id,omitempty"`
		UserId      int64              `bson:"user_id,omitempty"`
		GuestId     string             `bson:"guest_id,omitempty"`
		GuestName   string             `bson:"guestname,omitempty"`
		Platform    string             `bson:"platform,omitempty"`
		AndroidId   string             `bson:"android_id,omitempty"`
		AppVersion  string             `bson:"app_version,omitempty"`
		Country     string             `bson:"country,omitempty"`
		Imei        string             `bson:"imei,omitempty"`
		Channel     string             `bson:"channel,omitempty"`
		UserChannel string             `bson:"user_channel,omitempty"`
		CampaignId  string             `bson:"campaign_id,omitempty"`
		CreateTime  time.Time          `bson:"create_time,omitempty"`
		UpdateTime  time.Time          `bson:"update_time,omitempty"`
	}
)

func NewGuestModel(client *mongo.Client, globalConfig *config.Config, database, collection string) *GuestsModel {
	return &GuestsModel{
		globalConfig: globalConfig,
		client:       client,
		database:     database,
		collection:   collection,
	}
}

func (m *GuestsModel) WithCollection(f func(col *mongo.Collection, ctx context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection(m.collection)
	f(collection, ctx)
}

func (m *GuestsModel) Insert(data *Guests) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection(m.collection)
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	if res == nil {
		return primitive.ObjectID{}, db_errors.DBNilRes
	}
	return res.InsertedID.(primitive.ObjectID), err
}

func (m *GuestsModel) FindByGuestId(guestId string) (data *Guests, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		err = col.FindOne(ctx, bson.M{"guest_id": guestId}).Decode(&data)
	})
	return
}
