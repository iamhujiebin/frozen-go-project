package middleware

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"frozen-go-project/api/user-api/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
	"sort"
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


