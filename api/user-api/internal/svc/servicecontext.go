package svc

import (
	"frozen-go-project/api/user-api/internal/config"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"frozen-go-project/rpc/event-rpc/eventrpc"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"github.com/tal-tech/go-zero/rest/httpx"
	"github.com/tal-tech/go-zero/zrpc"
	"net/http"
	"strconv"
)

type ServiceContext struct {
	Config           config.Config
	UserRpc          userrpc.UserRpc
	BaseRpc          baserpc.BaseRpc
	EventRpc         eventrpc.EventRpc
	CheckAccessToken func(http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config:   c,
		UserRpc:  userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
		BaseRpc:  baserpc.NewBaseRpc(zrpc.MustNewClient(c.BaseRpc)),
		EventRpc: eventrpc.NewEventRpc(zrpc.MustNewClient(c.EventRpc)),
	}
	ctx.CheckAccessToken = CheckAccessToken(ctx)
	return ctx
}

func CheckAccessToken(ctx *ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseForm()
			userId, _ := strconv.Atoi(r.Form.Get("__user_id"))
			accessToken := r.Form.Get("accessToken")
			if userId > 0 {
				if len(accessToken) <= 0 {
					httpx.OkJson(w, resp_codes.CheckAccessTokenFail)
					return
				} else {
					_, err := ctx.UserRpc.CheckAccessToken(r.Context(), &userrpc.CheckAccessTokenReq{
						AccessToken: accessToken,
						UserInfo:    false,
					})
					if err != nil {
						httpx.OkJson(w, resp_codes.CheckAccessTokenFail)
						return
					}
				}
			}
			next(w, r)
		}
	}
}
