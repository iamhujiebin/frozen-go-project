package logic

import (
	"context"

	"frozen-go-project/rpc/test-rpc/internal/svc"
	"frozen-go-project/rpc/test-rpc/test_rpc"

	"github.com/tal-tech/go-zero/core/logx"
)

type TestTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestTestLogic {
	return &TestTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestTestLogic) TestTest(in *test_rpc.TestReq) (*test_rpc.TestRes, error) {
	// todo: add your logic here and delete this line

	return &test_rpc.TestRes{}, nil
}
