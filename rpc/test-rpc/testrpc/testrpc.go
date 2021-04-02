// Code generated by goctl. DO NOT EDIT!
// Source: test_rpc.proto

//go:generate mockgen -destination ./testrpc_mock.go -package testrpc -source $GOFILE

package testrpc

import (
	"context"

	"frozen-go-project/rpc/test-rpc/test_rpc"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	TestReq = test_rpc.TestReq
	TestRes = test_rpc.TestRes

	TestRpc interface {
		TestTest(ctx context.Context, in *TestReq) (*TestRes, error)
	}

	defaultTestRpc struct {
		cli zrpc.Client
	}
)

func NewTestRpc(cli zrpc.Client) TestRpc {
	return &defaultTestRpc{
		cli: cli,
	}
}

func (m *defaultTestRpc) TestTest(ctx context.Context, in *TestReq) (*TestRes, error) {
	client := test_rpc.NewTestRpcClient(m.cli.Conn())
	return client.TestTest(ctx, in)
}