package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserAssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAssetLogic {
	return &GetUserAssetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAssetLogic) GetUserAsset(in *user_rpc.GetUserAssetReq) (*user_rpc.GetUserAssetRes, error) {
	res, err := l.svcCtx.UserAssetMysqlModel.FindOneByUserId(in.UserId)
	if err != nil {
		return nil, err
	}
	pbUserAsset := new(user_rpc.UserAsset)
	_ = copier.Copy(pbUserAsset, res)
	return &user_rpc.GetUserAssetRes{
		UserAsset: pbUserAsset,
	}, nil
}
