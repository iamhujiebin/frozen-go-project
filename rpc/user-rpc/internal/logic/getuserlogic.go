package logic

import (
	"context"
	"frozen-go-project/common/codes/resp_codes"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/mongo"

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
	//todo 这里暂时只是支持loginname去找(即guest_id)
	user, err := l.svcCtx.UserMongoModel.FindByLoginName(in.LoginName)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return &user_rpc.GetUserRes{
			Code: resp_codes.ErrNotFound,
			User: nil,
		}, nil
	}
	pbUser := new(user_rpc.UserInfo)
	_ = copier.Copy(pbUser, user)
	pbUser.CreateTimeUnix = user.CreateTime.Unix()
	return &user_rpc.GetUserRes{
		User: pbUser,
	}, nil
}
