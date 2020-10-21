package logic

import (
	"context"
	"encoding/json"
	"flag"
	"frozen-go-project/common/errors/business_errors"
	"github.com/bitly/go-simplejson"
	"github.com/tal-tech/go-zero/core/conf"

	"frozen-go-project/rpc/base-rpc/internal/svc"
	base_rpc "frozen-go-project/rpc/base-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

var configFile = flag.String("p", "etc/pkg-config.yaml", "the config file")

type ConfigItems struct {
	Facebook struct {
		AppId  string
		AppKey string
	}
	Google struct {
		PayKey string
	}
	Adjust struct {
		AppToken           string
		PurchaseEventToken string
	}
	DefaultAvatar struct {
		Small string
		Big   string
	}
	FcmPush struct {
		ServerKey string
		SendId    string
	}
	ChatWs struct {
		Servers string
	}
}

type PkgConfig struct {
	Fatee    ConfigItems
	Supreshe ConfigItems
}

var PkgConfigJson *simplejson.Json

func init() {
	//load pkg configs
	var c PkgConfig
	conf.MustLoad(*configFile, &c)
	j, _ := json.Marshal(c)
	var err error
	PkgConfigJson, err = simplejson.NewJson(j)
	if err != nil {
		panic(err)
	}
}

type GetPkgConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPkgConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPkgConfigLogic {
	return &GetPkgConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPkgConfigLogic) GetPkgConfig(in *base_rpc.GetPkgConfigReq) (*base_rpc.GetPkgConfigResp, error) {
	value, err := PkgConfigJson.Get(in.PkgName).Get(in.Section).Get(in.Key).String()
	if err != nil {
		return nil, business_errors.NoConfigErr
	}
	return &base_rpc.GetPkgConfigResp{
		Value: value,
	}, nil
}
