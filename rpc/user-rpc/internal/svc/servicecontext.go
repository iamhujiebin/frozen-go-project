package svc

import (
	"context"
	"frozen-go-project/rpc/user-rpc/internal/config"
	mongoModel "frozen-go-project/rpc/user-rpc/internal/model/mongo"
	mysqlModel "frozen-go-project/rpc/user-rpc/internal/model/mysql"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/syncx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type ServiceContext struct {
	c                   config.Config
	MongoClient         *mongo.Client
	Cache               cache.Cache
	GuestMongoModel     *mongoModel.GuestsModel
	UserMongoModel      *mongoModel.UserModel
	UserExtMongoModel   *mongoModel.UserExtModel
	UserAssetMysqlModel *mysqlModel.UserAssetModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		c:                   c,
		UserAssetMysqlModel: mysqlModel.NewUserAssetModel(sqlx.NewMysql(c.DataSource), c.Cache, mysqlModel.TABLE_USER_ASSET),
	}
	if len(c.Mongo.Url) > 0 {
		initMongoModels(c, s)
	}
	return s
}

func initMongoModels(c config.Config, s *ServiceContext) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Mongo.OpTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetMaxPoolSize(c.Mongo.MaxPoolSize).SetRetryReads(true).
		ApplyURI(c.Mongo.Url))
	if err != nil {
		panic("connect to mongo fail:" + err.Error())
	}
	err = client.Ping(ctx, readpref.Nearest())
	if err != nil {
		panic("connect to mongo fail:" + err.Error())
	}
	s.MongoClient = client
	//syncx.NewSharedCalls 同一key同时取数据，只有一个协程取实际操作func
	s.Cache = cache.NewCache(c.Cache, syncx.NewSharedCalls(), cache.NewCacheStat("mongoc"), mongo.ErrNoDocuments,
		cache.WithExpiry(time.Second*time.Duration(c.CacheExpirySecond)))
	s.GuestMongoModel = mongoModel.NewGuestModel(client, &c, mongoModel.DB_FEWeb, mongoModel.COL_GUESTS)
	s.UserMongoModel = mongoModel.NewUserModel(client, s.Cache, &c, mongoModel.DB_FEWeb, mongoModel.COL_USERS)
	s.UserExtMongoModel = mongoModel.NewUserExtModel(client, &c, mongoModel.DB_FEWeb, mongoModel.COL_USEREXT)
}
