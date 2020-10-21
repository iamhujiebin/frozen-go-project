package mongo

import (
	"context"
	"frozen-go-project/common/errors/db_errors"
	"frozen-go-project/rpc/user-rpc/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type (
	UserModel struct {
		globalConfig *config.Config
		client       *mongo.Client
		database     string
		collection   string
	}

	User struct {
		Id          primitive.ObjectID `bson:"_id"`
		UserId      int64              `bson:"user_id"`
		AccessToken string             `bson:"access_token"`
		Avatar      string             `bson:"avatar"`
		LoginName   string             `bson:"loginname"`
		UserRole    string             `bson:"user_role"`
		NickName    string             `bson:"nick_name"`
		RegGuestId  string             `bson:"reg_guest_id"`
		Country     string             `bson:"country,omitempty"`
		Channel     string             `bson:"channel,omitempty"`
		UserChannel string             `bson:"user_channel,omitempty"`
		RegPkgName  string             `bson:"reg_pkg_name,omitempty"`
		CreateTime  time.Time          `bson:"create_time"`
		UpdateTime  time.Time          `bson:"update_time"`
	}
)

func NewUserModel(client *mongo.Client, globalConfig *config.Config, database, collection string) *UserModel {
	return &UserModel{
		globalConfig: globalConfig,
		client:       client,
		database:     database,
		collection:   collection,
	}
}

func (User) Database() string {
	return "FEWeb"
}

func (User) Collection() string {
	return "users"
}

func (m *UserModel) WithCollection(f func(col *mongo.Collection, ctx context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection(m.collection)
	f(collection, ctx)
}

func (m *UserModel) FindByLoginName(loginname string) (data *User, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		err = col.FindOne(ctx, bson.M{"loginname": loginname}).Decode(&data)
	})
	return
}

func (m *UserModel) InsertOne(user *User) (insertID primitive.ObjectID, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		res, iErr := col.InsertOne(ctx, user)
		if iErr != nil {
			err = iErr
		} else {
			insertID = res.InsertedID.(primitive.ObjectID)
		}
	})
	return
}

func (m *UserModel) UpdateOne(where, set bson.M) (err error) {
	if where == nil || set == nil {
		return db_errors.IllegalParams("where or set")
	}
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		update := bson.M{}
		if set != nil {
			update["$set"] = set
		}
		_, uRes := col.UpdateOne(ctx, where, update)
		if uRes != nil {
			err = uRes
		}
	})
	return
}
