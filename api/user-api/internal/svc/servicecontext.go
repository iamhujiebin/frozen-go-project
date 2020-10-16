package svc

import (
	"frozen-go-project/api/user-api/internal/config"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
