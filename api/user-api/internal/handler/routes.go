// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"frozen-go-project/api/user-api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/user/token",
			Handler: jwtHandler(serverCtx),
		},
		{
			Method:  http.MethodGet,
			Path:    "/guest/init",
			Handler: guestInitHandler(serverCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/guest/login",
			Handler: guestLoginHandler(serverCtx),
		},
	})

	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/user/info",
			Handler: getUserHandler(serverCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/user/add",
			Handler: addUserHandler(serverCtx),
		},
	}, rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret))
}
