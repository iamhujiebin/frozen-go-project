package main

import (
	"database/sql"
	"flag"
	"fmt"
	"frozen-go-project/models/commonconfig"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type Config struct {
	Mysql struct {
		Url             string
		MaxPoolSize     int
		ConnMaxLiveTime int
	}
	Cache cache.CacheConf
}

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	var c Config
	conf.MustLoad(*configFile, &c)

	db, err := sql.Open("mysql", c.Mysql.Url)
	if err != nil {
		panic("database/sql init fail " + c.Mysql.Url)
	}
	configModel := commonconfig.NewCommonConfigModel(sqlx.NewMysql(c.Mysql.Url), db, c.Cache)
	res, err := configModel.Insert(commonconfig.CommonConfig{
		Key:   "abc",
		Value: "efg",
	})
	fmt.Printf("res:%v,err:%v", res, err)
	res2, err := configModel.FindOne(1)
	fmt.Printf("res:%v,err:%v", res2.Value, err)
}
