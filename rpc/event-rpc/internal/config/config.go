package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogConf logx.LogConf
	Kafka   struct {
		Brokers       []string
		RequiredAcks  string
		RetryMax      int
		ReturnSuccess bool
		ReturnError   bool
	}
}
