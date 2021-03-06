// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"frozen-go-project/api/user-api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckAccessToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/userapi/user/token/:userId",
					Handler: jwtHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/userapi/:pkg/chat/dispatcher",
					Handler: dispatcherHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/userapi/anchor/recommend",
					Handler: anchorRecommendHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/userapi/user/action",
					Handler: userActionHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/userapi/guest/init",
				Handler: guestInitHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/userapi/guest/login",
				Handler: guestLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/userapi/hello",
				Handler: helloHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/userapi/user/info",
				Handler: getUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/userapi/user/add",
				Handler: addUserHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
