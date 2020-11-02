package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogConf           logx.LogConf
	Cache             cache.CacheConf
	CacheExpirySecond int64
	Mysql             struct {
		Url             string
		MaxPoolSize     int
		ConnMaxLiveTime int
	}
	Mongo struct {
		Url         string
		MaxPoolSize uint64
		OpTimeout   uint64
	}
}
