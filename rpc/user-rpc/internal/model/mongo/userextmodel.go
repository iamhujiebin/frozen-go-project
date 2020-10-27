package mongo

import (
	"context"
	"frozen-go-project/rpc/user-rpc/internal/config"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type (
	UserExtModel struct {
		globalConfig *config.Config
		client       *mongo.Client
		database     string
		collection   string
	}

	UserExt struct {
		Id          primitive.ObjectID `bson:"_id,omitempty"`
		UserId      int64              `bson:"user_id"`
		ActionPoint struct {
			PayChannel    int64 `bson:"pay_channel"`
			VipPromotion  int64 `bson:"vip_promotion"`
			CoinPromotion int64 `bson:"coin_promotion"`
		} `bson:"action_point"`
		CreateTime time.Time `bson:"create_time"`
		UpdateTime time.Time `bson:"update_time"`
	}
)

func NewUserExtModel(client *mongo.Client, globalConfig *config.Config, database, collection string) *UserExtModel {
	return &UserExtModel{
		globalConfig: globalConfig,
		client:       client,
		database:     database,
		collection:   collection,
	}
}

func (m *UserExtModel) WithCollection(f func(col *mongo.Collection, ctx context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection(m.collection)
	f(collection, ctx)
}

func (m *UserExtModel) WithTransaction(f func(sessCtx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := m.client.StartSession(opts)
	if err != nil {
		return nil, err
	}
	defer sess.EndSession(ctx)
	txnOpts := options.Transaction().SetReadPreference(readpref.PrimaryPreferred()) //mongo的一些mode：nearest,secondary,primary等等
	return sess.WithTransaction(ctx, f, txnOpts)
}

func (m *UserExtModel) UpsertUserExt(where, set, inc bson.M) (userExt *UserExt, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		update := bson.M{}
		if len(set) > 0 {
			update["$set"] = set
		}
		if len(inc) > 0 {
			update["$inc"] = inc
		}
		err = col.FindOneAndUpdate(ctx, where, update, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)).Decode(&userExt)
	})
	return
}
