package logic

import (
	"context"
	"frozen-go-project/rpc/user-rpc/userrpc"

	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GuestInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuestInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) GuestInitLogic {
	return GuestInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuestInitLogic) GuestInit(req types.GuestInitRequest) (*types.CommonResponse, error) {
	res, err := l.svcCtx.UserRpc.GuestInit(l.ctx, &userrpc.GuestInitReq{
		GuestId:     req.GuestId,
		Platform:    req.Platform,
		AndroidId:   req.AndroidId,
		AppVersion:  req.AppVersion,
		Country:     req.Country,
		Imei:        req.Imei,
		Channel:     req.Channel,
		CampaignId:  req.CampaignId,
		UserChannel: req.UserChannel,
	})
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    0,
		Body:    res.Guest,
		Message: "todo",
	}, nil
}
