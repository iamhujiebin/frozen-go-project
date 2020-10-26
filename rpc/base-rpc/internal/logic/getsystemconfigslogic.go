package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"frozen-go-project/rpc/base-rpc/internal/svc"

	base_rpc "frozen-go-project/rpc/base-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSystemConfigsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSystemConfigsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemConfigsLogic {
	return &GetSystemConfigsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSystemConfigsLogic) GetSystemConfigs(in *base_rpc.GetSystemConfigReq) (*base_rpc.GetSystemConfigRes, error) {
	res, err := l.svcCtx.CommonConfigModel.FindSystemConfigs(in.Section, in.Keys)
	if err != nil {
		return nil, err
	}
	var pbSystemConfigs []*base_rpc.SystemConfig
	for k := range res {
		pbSystemConfig := new(base_rpc.SystemConfig)
		_ = copier.Copy(pbSystemConfig, res[k])
		pbSystemConfigs = append(pbSystemConfigs, pbSystemConfig)
	}
	return &base_rpc.GetSystemConfigRes{
		Items: pbSystemConfigs,
	}, nil
}
