package logic

import (
	"context"
	"frozen-go-project/common/public_method"
	mongoModel "frozen-go-project/rpc/user-rpc/internal/model/mongo"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"time"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GuestInitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGuestInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuestInitLogic {
	return &GuestInitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GuestInitLogic) GuestInit(in *user_rpc.GuestInitReq) (*user_rpc.GuestInitRes, error) {
	pbGuest := new(user_rpc.GuestInfo)
	if len(in.GuestId) > 0 {
		guest, err := l.svcCtx.GuestMongoModel.FindByGuestId(in.GuestId)
		if guest != nil && err == nil {
			_ = copier.Copy(pbGuest, guest)
			return &user_rpc.GuestInitRes{Guest: pbGuest}, nil
		}
	}
	guest := &mongoModel.Guests{
		GuestId:     uuid.New().String(),
		GuestName:   "guest" + public_method.GetRandomCodeFromNumber(5),
		Platform:    in.Platform,
		AndroidId:   in.AndroidId,
		AppVersion:  in.AppVersion,
		Country:     in.Country,
		Imei:        in.Imei,
		Channel:     in.Channel,
		UserChannel: in.UserChannel,
		CampaignId:  in.CampaignId,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	_, err := l.svcCtx.GuestMongoModel.Insert(guest)
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(pbGuest, guest)
	return &user_rpc.GuestInitRes{
		Guest: pbGuest,
	}, nil
}
