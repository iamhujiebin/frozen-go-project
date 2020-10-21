package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user_rpc.AddUserReq) (*user_rpc.AddUserRes, error) {
	user, err := l.svcCtx.UserMysqlModel.AddUserTx(in.Avatar)
	if err != nil {
		return nil, err
	}
	pbUser := new(user_rpc.UserInfo)
	copier.Copy(pbUser, user)
	pbUser.CreateTimeUnix = user.CreateTime.Unix()
	return &user_rpc.AddUserRes{
		User: pbUser,
	}, nil
}
