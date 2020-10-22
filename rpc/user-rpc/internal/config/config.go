package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogConf    logx.LogConf
	DataSource string          // 手动代码
	Cache      cache.CacheConf // 手动代码
	Mongo      struct {
		Url         string
		MaxPoolSize uint64
		OpTimeout   uint64
	}
}
