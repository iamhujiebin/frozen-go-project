package logic

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/jinzhu/copier"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddActionPointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActionPointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActionPointLogic {
	return &AddActionPointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActionPointLogic) AddActionPoint(in *user_rpc.AddActionPointReq) (*user_rpc.AddActionPointRes, error) {
	where := bson.M{"user_id": in.UserId}
	inc := bson.M{
		"pay_channel":    in.PayChannelPoint,
		"vip_promotion":  in.VipPromotionPoint,
		"coin_promotion": in.CoinPromotionPoint,
	}
	res, err := l.svcCtx.UserExtMongoModel.UpsertUserExt(where, nil, inc)
	if err != nil {
		return nil, err
	}
	pbUserExt := new(user_rpc.UserExt)
	_ = copier.Copy(pbUserExt, res)
	return &user_rpc.AddActionPointRes{
		UserExt: pbUserExt,
	}, nil
}
