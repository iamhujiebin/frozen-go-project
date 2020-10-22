package svc

import (
	"frozen-go-project/api/user-api/internal/config"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userrpc.UserRpc
	BaseRpc baserpc.BaseRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.LogConf)
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
		BaseRpc: baserpc.NewBaseRpc(zrpc.MustNewClient(c.BaseRpc)),
	}
}
