package middleware

import (
	"context"
	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/common/enum"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

func CheckBan(ctx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if len(r.Form.Get(enum.CommonParams.UserId)) > 0 || len(r.Form.Get(enum.CommonParams.GuestId)) > 0 {
				userId, _ := strconv.Atoi(r.Form.Get(enum.CommonParams.UserId))
				guestId := r.Form.Get(enum.CommonParams.GuestId)
				res, _ := ctx.BaseRpc.IsBan(context.Background(), &baserpc.IsBanReq{
					UserId:  int64(userId),
					GuestId: guestId,
				})
				if res != nil && res.IsBan {
					httpx.OkJson(w, resp_codes.BanFail)
					return
				}
			}
			next(w, r)
		}
	}
}
