package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogConf logx.LogConf
	Mongo   struct {
		Url         string
		MaxPoolSize uint64
		OpTimeout   uint64
	}
}
