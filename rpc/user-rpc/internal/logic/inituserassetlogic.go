package logic

import (
	"context"
	mysqlModel "frozen-go-project/rpc/user-rpc/internal/model/mysql"
	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"
	"github.com/jinzhu/copier"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type InitUserAssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitUserAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitUserAssetLogic {
	return &InitUserAssetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitUserAssetLogic) InitUserAsset(in *user_rpc.InitUserAssetReq) (*user_rpc.InitUserAssetRes, error) {
	userAsset := new(mysqlModel.UserAsset)
	_ = copier.Copy(userAsset, in.UserAsset)
	userAsset.CreateTime, userAsset.UpdateTime = time.Now(), time.Now()
	userAsset.VipEffectEnd = time.Unix(0, 0)
	//res, err := l.svcCtx.UserAssetMysqlModel.Insert(*userAsset)
	res, err := l.svcCtx.UserAssetMysqlModel.InsertByGorm(userAsset)
	if err != nil {
		return nil, err
	}
	pbUserAsset := new(user_rpc.UserAsset)
	_ = copier.Copy(pbUserAsset, res)
	return &user_rpc.InitUserAssetRes{
		UserAsset: pbUserAsset,
	}, nil
}
