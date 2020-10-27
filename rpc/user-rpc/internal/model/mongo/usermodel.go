package mongo

import (
	"context"
	"frozen-go-project/common/enum"
	"frozen-go-project/common/errors/db_errors"
	"frozen-go-project/common/public_method"
	"frozen-go-project/rpc/user-rpc/internal/config"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
		Id          primitive.ObjectID `bson:"_id,omitempty"`
		UserId      int64              `bson:"user_id"`
		AccessToken string             `bson:"access_token"`
		Avatar      string             `bson:"avatar"`
		LoginName   string             `bson:"loginname"`
		UserRole    enum.UserRole      `bson:"user_role"`
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

func (m *UserModel) WithCollection(f func(col *mongo.Collection, ctx context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection(m.collection)
	f(collection, ctx)
}

func (m *UserModel) WithTransaction(f func(sessCtx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
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

func (m *UserModel) FindByLoginName(loginname string) (data *User, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		err = col.FindOne(ctx, bson.M{"loginname": loginname}).Decode(&data)
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

func (m *UserModel) GuestRegister(in *user_rpc.GuestLoginReq) (interface{}, error) {
	return m.WithTransaction(func(sessCtx mongo.SessionContext) (interface{}, error) {
		//获取自增用户id
		idCol := m.client.Database(DB_FEWeb).Collection(COL_IDS)
		idTable := new(Ids)
		err := idCol.FindOneAndUpdate(sessCtx, bson.M{"tablename": COL_USERS}, bson.M{"$inc": bson.M{"id": 1}},
			options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&idTable)
		if err != nil {
			return nil, err
		}
		if idTable.Id <= 0 {
			return nil, db_errors.DBNilRes
		}
		user := &User{
			UserId:      idTable.Id,
			AccessToken: uuid.New().String(),
			Avatar:      "http://some.default.avatar.png",
			LoginName:   in.GuestId,
			UserRole:    enum.UserRoleEnum.Normal,
			NickName:    "frozen_" + public_method.GetRandomString(5),
			RegGuestId:  in.GuestId,
			Country:     in.Country,
			Channel:     in.Channel,
			UserChannel: in.UserChannel,
			RegPkgName:  in.PkgName,
			CreateTime:  time.Time{},
			UpdateTime:  time.Time{},
		}
		userCol := m.client.Database(DB_FEWeb).Collection(COL_USERS)
		_, err = userCol.InsertOne(sessCtx, user)
		if err != nil {
			return nil, err
		}
		guestCol := m.client.Database(DB_FEWeb).Collection(COL_GUESTS)
		guest := new(Guests)
		guest.CreateTime = time.Now()
		guest.UpdateTime = time.Now()
		guest.GuestName = "guest" + public_method.GetRandomCodeFromNumber(5)
		_ = copier.Copy(guest, in)
		guest.UserId = user.UserId
		err = guestCol.FindOneAndUpdate(sessCtx, bson.M{"guest_id": in.GuestId}, bson.M{"$set": guest}, //此处直接$set:guest，如果guest中字段有omitempty，那么如果不赋值,就不会更新对应字段!!!
			options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)).Err()
		if err != nil {
			return nil, err
		}
		return user, nil
	})
}

func (m *UserModel) PageUsers(where bson.M, skip, limit int64) (users []*User, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		cur, cErr := col.Find(ctx, where, options.Find().SetSort(bson.M{"anchor_weight": -1, "active_at": -1}).SetSkip(skip).SetLimit(limit))
		if cErr != nil {
			err = cErr
			return
		}
		for cur.Next(ctx) {
			var user *User
			err = cur.Decode(&user)
			if err != nil {
				return
			}
			users = append(users, user)
		}
	})
	return
}

func (m *UserModel) FindOneByAccessToken(accessToken string) (user *User, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		err = col.FindOne(ctx, bson.M{"accessToken": accessToken}).Decode(&user)
	})
	return
}
