package logic

import (
	"context"
	"frozen-go-project/api/user-api/internal/dto"
	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"
)

type GuestLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuestLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) GuestLoginLogic {
	return GuestLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuestLoginLogic) GuestLogin(req types.GuestLoginRequest) (*types.CommonResponse, error) {
	response := new(types.CommonResponse)
	res, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userrpc.GetUserReq{
		LoginName: req.GuestId,
	})
	if err != nil {
		response.Code = resp_codes.RpcError
		response.Message = err.Error()
		return response, nil
	}
	//注册
	if res != nil && res.Code == resp_codes.ErrNotFound {
		respBody, err := l.doUserRegister(&req)
		if err != nil {
			response.Code = resp_codes.RpcError
			response.Message = err.Error()
			return response, nil
		}
		response.Body = respBody
	}
	//登录
	if res != nil && res.User != nil && res.User.UserId > 0 {
		respBody := new(dto.GuestLoginResponse)
		_ = copier.Copy(respBody, res.User)
		userAsset, err := l.doGuestLogin(res.User.UserId)
		if err != nil {
			response.Code = resp_codes.RpcError
			response.Message = err.Error()
		}
		respBody.AvailableCoin = userAsset.AvailableCoin
		response.Body = respBody
	}

	return response, nil
}

func (l *GuestLoginLogic) doUserRegister(req *types.GuestLoginRequest) (*dto.GuestLoginResponse, error) {
	//todo 可以考虑都骄傲给rpc做mongo+mysql的联合事务
	addUserRes, err := l.svcCtx.UserRpc.AddUser(l.ctx, &userrpc.AddUserReq{
		Avatar:      "http://default_avatr.png",
		GuestId:     req.GuestId,
		PkgName:     req.PkgName,
		Channel:     req.PkgChannel,
		UserChannel: req.UserChannel,
		Platform:    req.Platform,
		Country:     req.Country,
	})
	if err != nil {
		return nil, err
	}
	respBody := new(dto.GuestLoginResponse)
	_ = copier.Copy(respBody, addUserRes.User)
	//todo 这里写mysql会比较慢，所以不设置deadline
	_, err = l.svcCtx.UserRpc.InitUserAsset(context.Background(), &userrpc.InitUserAssetReq{UserAsset: &userrpc.UserAsset{
		UserId:                respBody.UserId,
		AvailableCoin:         0,
		AccumulatedCoin:       0,
		FreeChatTimes:         0,
		FreeCallTimes:         0,
		AvailableSilverCoin:   0,
		AccumulatedSilverCoin: 0,
		VipEffectEndMs:        0,
	}})
	if err != nil {
		return nil, err
	} else {
		respBody.AvailableCoin = 0
	}
	return respBody, nil
}

func (l *GuestLoginLogic) doGuestLogin(userId int64) (*userrpc.UserAsset, error) {
	userAssetRes, err := l.svcCtx.UserRpc.GetUserAsset(l.ctx, &userrpc.GetUserAssetReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return userAssetRes.UserAsset, nil
}
