package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogConf           logx.LogConf
	DataSource        string
	Cache             cache.CacheConf
	CacheExpirySecond int64
	Mongo             struct {
		Url         string
		MaxPoolSize uint64
		OpTimeout   uint64
	}
}
