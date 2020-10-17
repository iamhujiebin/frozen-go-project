package logic

import (
	"context"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"strconv"

	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req types.AddUserRequest) (*types.AddUserResponse, error) {
	user, err := l.svcCtx.UserRpc.AddUser(l.ctx, &userrpc.AddUserReq{Avatar: req.Avatar})
	if err != nil {
		return nil, err
	}
	return &types.AddUserResponse{
		UserId:         strconv.Itoa(int(user.User.UserId)),
		Avatar:         user.User.Avatar,
		AccessToken:    user.User.AccessToken,
		CreateTimeUnix: user.User.CreateTimeUnix,
	}, nil
}
