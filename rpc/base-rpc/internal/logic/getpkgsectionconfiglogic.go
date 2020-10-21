package logic

import (
	"context"

	"frozen-go-project/rpc/base-rpc/internal/svc"
	base_rpc "frozen-go-project/rpc/base-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetPkgSectionConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPkgSectionConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPkgSectionConfigLogic {
	return &GetPkgSectionConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPkgSectionConfigLogic) GetPkgSectionConfig(in *base_rpc.GetPkgSectionConfigReq) (*base_rpc.GetPkgSectionConfigResp, error) {
	// todo: add your logic here and delete this line

	return &base_rpc.GetPkgSectionConfigResp{}, nil
}
