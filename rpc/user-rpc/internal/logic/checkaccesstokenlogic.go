package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckAccessTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAccessTokenLogic {
	return &CheckAccessTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAccessTokenLogic) CheckAccessToken(in *user_rpc.CheckAccessTokenReq) (*user_rpc.CheckAccessTokenRes, error) {
	res, err := l.svcCtx.UserMongoModel.FindOneByAccessToken(in.AccessToken)
	if err != nil {
		return nil, err
	}
	pbUser := new(user_rpc.UserInfo)
	_ = copier.Copy(pbUser, res)
	if in.UserInfo {
		return &user_rpc.CheckAccessTokenRes{
			Success:  true,
			UserInfo: pbUser,
		}, nil
	} else {
		return &user_rpc.CheckAccessTokenRes{
			Success: true,
		}, nil
	}
}
