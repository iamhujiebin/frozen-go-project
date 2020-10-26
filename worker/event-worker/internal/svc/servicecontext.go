package svc

import (
	"frozen-go-project/rpc/base-rpc/baserpc"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	BaseRpc baserpc.BaseRpc
}

var ServiceCtx *ServiceContext

func InitServiceContext(baseRpc zrpc.RpcClientConf) {
	ServiceCtx = &ServiceContext{
		BaseRpc: baserpc.NewBaseRpc(zrpc.MustNewClient(baseRpc)),
	}
}
