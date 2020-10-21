package logic

import (
	"context"
	"frozen-go-project/common/errors/business_errors"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GuestLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGuestLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuestLoginLogic {
	return &GuestLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GuestLoginLogic) GuestLogin(in *user_rpc.GuestLoginReq) (*user_rpc.GuestLoginRes, error) {
	if len(in.GuestId) <= 0 {
		return nil, business_errors.EmptyParams("guest_id")
	}
	user, err := l.svcCtx.UserMongoModel.FindByLoginName(in.GuestId)
	//登录
	if err == nil && user != nil {
		newAccessToken := uuid.New().String()
		set := bson.M{
			"accessToken": newAccessToken,
			"active_at":   time.Now(),
		}
		err = l.svcCtx.UserMongoModel.UpdateOne(bson.M{"user_id": user.UserId}, set)
		if err != nil {
			return nil, err
		}
		pbUser := new(user_rpc.UserInfo)
		user.AccessToken = newAccessToken
		_ = copier.Copy(pbUser, user)
		return &user_rpc.GuestLoginRes{
			User: pbUser,
		}, nil
	}
	//注册-事务操作
	newUser, err := l.svcCtx.UserMongoModel.GuestRegister(in)
	if err != nil {
		return nil, err
	}
	pbUser := new(user_rpc.UserInfo)
	_ = copier.Copy(pbUser, newUser)
	return &user_rpc.GuestLoginRes{
		User: pbUser,
	}, nil
}
