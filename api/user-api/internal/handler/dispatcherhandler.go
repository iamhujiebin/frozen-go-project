package handler

import (
	"net/http"

	"frozen-go-project/api/user-api/internal/logic"
	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func dispatcherHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DispatcherRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDispatcherLogic(r.Context(), ctx)
		resp, err := l.Dispatcher(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
