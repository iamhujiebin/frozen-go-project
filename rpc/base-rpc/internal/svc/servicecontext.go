package svc

import (
	"context"
	"frozen-go-project/rpc/base-rpc/internal/config"
	mongoModel "frozen-go-project/rpc/base-rpc/internal/model/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type ServiceContext struct {
	c                 config.Config
	MongoClient       *mongo.Client
	CommonConfigModel *mongoModel.SystemConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		c: c,
	}
	if len(c.Mongo.Url) > 0 {
		initMongoModels(c, s)
	}
	return s
}

func initMongoModels(c config.Config, s *ServiceContext) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Mongo.OpTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetMaxPoolSize(c.Mongo.MaxPoolSize).ApplyURI(c.Mongo.Url))
	if err != nil {
		panic("connect to mongo fail:" + err.Error())
	}
	err = client.Ping(ctx, readpref.Nearest())
	if err != nil {
		panic("connect to mongo fail:" + err.Error())
	}
	s.MongoClient = client
	s.CommonConfigModel = mongoModel.NewCommonConfigModel(client, &c, mongoModel.DB_FEWeb, mongoModel.COL_COMMON_CONFIG)
}
