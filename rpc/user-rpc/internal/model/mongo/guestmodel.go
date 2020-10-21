package mongo

import (
	"context"
	"frozen-go-project/common"
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
		UserId      int                `bson:"user_id,omitempty"`
		GuestId     string             `bson:"guest_id"`
		GuestName   string             `bson:"guestname"`
		Platform    string             `bson:"platform"`
		AndroidId   string             `bson:"android_id"`
		AppVersion  string             `bson:"app_version"`
		Country     string             `bson:"country"`
		Imei        string             `bson:"imei"`
		Channel     string             `bson:"channel"`
		UserChannel string             `bson:"user_channel"`
		CampaignId  string             `bson:"campaign_id"`
		CreateTime  time.Time          `bson:"create_time"`
		UpdateTime  time.Time          `bson:"update_time"`
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

func (Guests) Database() string {
	return "FEWeb"
}

func (Guests) Collection() string {
	return "guests"
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
		return primitive.ObjectID{}, common.DBNilRes
	}
	return res.InsertedID.(primitive.ObjectID), err
}

func (m *GuestsModel) FindByGuestId(guestId string) (data *Guests, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		err = col.FindOne(ctx, bson.M{"guest_id": guestId}).Decode(&data)
	})
	return
}
