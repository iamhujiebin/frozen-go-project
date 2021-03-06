package svc

import (
	"context"
	"frozen-go-project/rpc/user-rpc/internal/config"
	mongoModel "frozen-go-project/rpc/user-rpc/internal/model/mongo"
	mysqlModel "frozen-go-project/rpc/user-rpc/internal/model/mysql"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/syncx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type ServiceContext struct {
	c                   config.Config
	MongoClient         *mongo.Client
	MysqlClient         *gorm.DB
	Cache               cache.Cache
	GuestMongoModel     *mongoModel.GuestsModel
	UserMongoModel      *mongoModel.UserModel
	UserExtMongoModel   *mongoModel.UserExtModel
	UserAssetMysqlModel *mysqlModel.UserAssetModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.Url), &gorm.Config{})
	if err != nil {
		logx.Errorf("gorm connect to db fail :%s", err.Error())
	} else {
		rawDb, err := db.DB()
		if err == nil {
			rawDb.SetConnMaxLifetime(time.Second * time.Duration(c.Mysql.ConnMaxLiveTime))
			rawDb.SetMaxOpenConns(c.Mysql.MaxPoolSize)
			rawDb.SetMaxIdleConns(c.Mysql.MaxPoolSize)
		}
	}
	s := &ServiceContext{
		c:                   c,
		MysqlClient:         db,
		UserAssetMysqlModel: mysqlModel.NewUserAssetModel(sqlx.NewMysql(c.Mysql.Url), db, c.Cache, mysqlModel.TABLE_USER_ASSET),
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
