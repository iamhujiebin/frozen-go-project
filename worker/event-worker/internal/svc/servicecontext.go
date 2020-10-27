package svc

import (
	"frozen-go-project/rpc/base-rpc/baserpc"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	BaseRpc baserpc.BaseRpc
	UserRpc userrpc.UserRpc
}

var ServiceCtx *ServiceContext

func InitServiceContext(baseRpc, userRpc zrpc.RpcClientConf) {
	ServiceCtx = &ServiceContext{
		BaseRpc: baserpc.NewBaseRpc(zrpc.MustNewClient(baseRpc)),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(userRpc)),
	}
}
