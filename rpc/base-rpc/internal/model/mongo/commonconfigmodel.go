package mongo

import (
	"context"
	"frozen-go-project/common/system_config"
	"frozen-go-project/rpc/base-rpc/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type (
	SystemConfigModel struct {
		globalConfig *config.Config
		client       *mongo.Client
		database     string
		collection   string
	}

	SystemConfig struct {
		Id       primitive.ObjectID          `bson:"_id,omitempty"`
		Section  system_config.ConfigSection `bson:"section"`
		Key      system_config.ConfigKey     `bson:"key"`
		Category string                      `bson:"category"`
		Value    string                      `bson:"value"`
	}
)

func NewCommonConfigModel(client *mongo.Client, globalConfig *config.Config, database, collection string) *SystemConfigModel {
	return &SystemConfigModel{
		globalConfig: globalConfig,
		client:       client,
		database:     database,
		collection:   collection,
	}
}

func (m *SystemConfigModel) WithCollection(f func(col *mongo.Collection, ctx context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.globalConfig.Mongo.OpTimeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection(m.collection)
	f(collection, ctx)
}

func (m *SystemConfigModel) FindSystemConfigs(section string, keys []string) (data []*SystemConfig, err error) {
	m.WithCollection(func(col *mongo.Collection, ctx context.Context) {
		where := bson.M{
			"section": section,
			"key":     bson.M{"$in": keys},
		}
		if len(keys) > 0 {
			where["key"] = bson.M{"$in": keys}
		}
		cur, cErr := col.Find(ctx, where)
		if cErr != nil {
			err = cErr
			return
		}
		for cur.Next(ctx) {
			var systemConfig *SystemConfig
			err = cur.Decode(&systemConfig)
			if err != nil {
				return
			}
			data = append(data, systemConfig)
		}
	})
	return
}
