package svc

import (
	"frozen-go-project/rpc/base-rpc/internal/config"
)

type ServiceContext struct {
	c config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
	}
}
