package middleware

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/common/enum"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"sort"
	"strconv"
)

func CheckSign(ctx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseForm()
			paramsKeyList := make([]string, 0, len(r.Form))
			var signValue string
			for k := range r.Form {
				if k != "sign" {
					paramsKeyList = append(paramsKeyList, k)
				} else {
					signValue = r.Form.Get("sign")
				}
			}
			if len(signValue) == 0 {
				logx.Errorf("should have check sign")
				//httpx.OkJson(w, resp_codes.CheckSignFail)
				//return
			}
			sort.Strings(paramsKeyList)
			var buffer bytes.Buffer
			buffer.WriteString(ctx.Config.SignSecretKey)
			for _, k := range paramsKeyList {
				buffer.WriteString(r.Form.Get(k))
			}
			h := md5.New()
			h.Write(buffer.Bytes())
			expectSign := hex.EncodeToString(h.Sum(nil))
			if expectSign != signValue {
				logx.Errorf("should have check sign")
				//httpx.OkJson(w, resp_codes.CheckSignFail)
				//return
			}
			next(w, r)
		}
	}
}

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
