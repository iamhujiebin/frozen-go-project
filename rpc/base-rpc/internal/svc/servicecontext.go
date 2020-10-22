package svc

import (
	"frozen-go-project/rpc/base-rpc/internal/config"
	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceContext struct {
	c config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.LogConf)
	return &ServiceContext{
		c: c,
	}
}
