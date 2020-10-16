package logic

import (
	"context"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"strconv"

	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.GetUserRequest) (*types.GetUserResponse, error) {
	userId, _ := strconv.Atoi(req.UserId)
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userrpc.GetUserReq{UserId: int64(userId)})
	if err != nil {
		return nil, err
	}
	return &types.GetUserResponse{
		UserId:         strconv.Itoa(int(user.User.UserId)),
		AccessToken:    user.User.AccessToken,
		Avatar:         user.User.Avatar,
		CreateTimeUnix: user.User.CreateTimeUnix,
	}, nil
}
