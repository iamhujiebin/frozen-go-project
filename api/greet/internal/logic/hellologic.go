package logic

import (
	"context"

	"frozen-go-project/api/greet/internal/svc"
	"frozen-go-project/api/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) HelloLogic {
	return HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req types.CommonParams) (*types.CommonResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CommonResponse{}, nil
}
