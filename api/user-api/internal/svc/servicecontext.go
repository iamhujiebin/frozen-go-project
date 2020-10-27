package svc

import (
	"frozen-go-project/api/user-api/internal/config"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"frozen-go-project/rpc/event-rpc/eventrpc"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  userrpc.UserRpc
	BaseRpc  baserpc.BaseRpc
	EventRpc eventrpc.EventRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRpc:  userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
		BaseRpc:  baserpc.NewBaseRpc(zrpc.MustNewClient(c.BaseRpc)),
		EventRpc: eventrpc.NewEventRpc(zrpc.MustNewClient(c.EventRpc)),
	}
}
