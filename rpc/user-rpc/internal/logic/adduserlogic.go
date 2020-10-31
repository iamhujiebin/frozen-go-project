package logic

import (
	"context"
	"frozen-go-project/common/errors/business_errors"
	mongoModel "frozen-go-project/rpc/user-rpc/internal/model/mongo"
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
	if len(in.GuestId) <= 0 {
		return nil, business_errors.EmptyParams("guest_id")
	}
	user, err := l.svcCtx.UserMongoModel.AddUser(in)
	if err != nil {
		return nil, err
	}
	pbUser := new(user_rpc.UserInfo)
	_ = copier.Copy(pbUser, user)
	pbUser.CreateTimeUnix = user.(*mongoModel.User).CreateTime.Unix()
	return &user_rpc.AddUserRes{
		User: pbUser,
	}, nil
}
