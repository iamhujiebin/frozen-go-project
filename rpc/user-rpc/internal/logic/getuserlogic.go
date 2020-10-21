package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user_rpc.GetUserReq) (*user_rpc.GetUserRes, error) {
	user, err := l.svcCtx.UserMysqlModel.FindOne(in.UserId)
	if err != nil {
		return nil, err
	}
	pbUser := new(user_rpc.UserInfo)
	_ = copier.Copy(pbUser, user)
	pbUser.CreateTimeUnix = user.CreateTime.Unix()
	return &user_rpc.GetUserRes{
		User: pbUser,
	}, nil
}
