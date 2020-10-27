package middlerware

import (
	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

func CheckAccessToken(ctx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
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
