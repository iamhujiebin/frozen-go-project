package logic

import (
	"context"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/rpc/user-rpc/userrpc"

	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GuestLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuestLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) GuestLoginLogic {
	return GuestLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuestLoginLogic) GuestLogin(req types.GuestLoginRequest) (*types.CommonResponse, error) {
	response := new(types.CommonResponse)
	res, err := l.svcCtx.UserRpc.GuestLogin(l.ctx, &userrpc.GuestLoginReq{
		GuestId:     req.GuestId,
		Platform:    req.Platform,
		Country:     req.Country,
		Channel:     req.PkgChannel,
		UserChannel: req.UserChannel,
		PkgName:     req.PkgName,
	})
	if err != nil {
		response.Code = resp_codes.RpcError
		response.Message = err.Error()
	} else {
		response.Body = res
	}
	return response, nil
}
