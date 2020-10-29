package logic

import (
	"context"

	"frozen-go-project/rpc/base-rpc/internal/svc"
	base_rpc "frozen-go-project/rpc/base-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsBanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsBanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsBanLogic {
	return &IsBanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsBanLogic) IsBan(in *base_rpc.IsBanReq) (*base_rpc.IsBanRes, error) {
	ban, _ := l.svcCtx.BanModel.FindOne(in.UserId, in.GuestId)
	if ban != nil {
		return &base_rpc.IsBanRes{
			IsBan: true,
		}, nil
	}
	return &base_rpc.IsBanRes{
		IsBan: false,
	}, nil
}
