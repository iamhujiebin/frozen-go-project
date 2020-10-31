// Code generated by goctl. DO NOT EDIT!
// Source: user_rpc.proto

//go:generate mockgen -destination ./userrpc_mock.go -package userrpc -source $GOFILE

package userrpc

import (
	"context"

	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/jsonx"
	"github.com/tal-tech/go-zero/zrpc"
)

type (
	UserRpc interface {
		GetUser(ctx context.Context, in *GetUserReq) (*GetUserRes, error)
		AddUser(ctx context.Context, in *AddUserReq) (*AddUserRes, error)
		GuestInit(ctx context.Context, in *GuestInitReq) (*GuestInitRes, error)
		PageAnchorRecommend(ctx context.Context, in *PageAnchorRecommendReq) (*PageAnchorRecommendRes, error)
		AddActionPoint(ctx context.Context, in *AddActionPointReq) (*AddActionPointRes, error)
		CheckAccessToken(ctx context.Context, in *CheckAccessTokenReq) (*CheckAccessTokenRes, error)
		GetUserAsset(ctx context.Context, in *GetUserAssetReq) (*GetUserAssetRes, error)
		InitUserAsset(ctx context.Context, in *InitUserAssetReq) (*InitUserAssetRes, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

func (m *defaultUserRpc) GetUser(ctx context.Context, in *GetUserReq) (*GetUserRes, error) {
	var request user_rpc.GetUserReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.GetUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret GetUserRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultUserRpc) AddUser(ctx context.Context, in *AddUserReq) (*AddUserRes, error) {
	var request user_rpc.AddUserReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.AddUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret AddUserRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultUserRpc) GuestInit(ctx context.Context, in *GuestInitReq) (*GuestInitRes, error) {
	var request user_rpc.GuestInitReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.GuestInit(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret GuestInitRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultUserRpc) PageAnchorRecommend(ctx context.Context, in *PageAnchorRecommendReq) (*PageAnchorRecommendRes, error) {
	var request user_rpc.PageAnchorRecommendReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.PageAnchorRecommend(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret PageAnchorRecommendRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultUserRpc) AddActionPoint(ctx context.Context, in *AddActionPointReq) (*AddActionPointRes, error) {
	var request user_rpc.AddActionPointReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.AddActionPoint(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret AddActionPointRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultUserRpc) CheckAccessToken(ctx context.Context, in *CheckAccessTokenReq) (*CheckAccessTokenRes, error) {
	var request user_rpc.CheckAccessTokenReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.CheckAccessToken(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret CheckAccessTokenRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultUserRpc) GetUserAsset(ctx context.Context, in *GetUserAssetReq) (*GetUserAssetRes, error) {
	var request user_rpc.GetUserAssetReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.GetUserAsset(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret GetUserAssetRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultUserRpc) InitUserAsset(ctx context.Context, in *InitUserAssetReq) (*InitUserAssetRes, error) {
	var request user_rpc.InitUserAssetReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := user_rpc.NewUserRpcClient(m.cli.Conn())
	resp, err := client.InitUserAsset(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret InitUserAssetRes
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}
