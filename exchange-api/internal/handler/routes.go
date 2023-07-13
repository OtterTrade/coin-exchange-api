// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	admin "github.com/otter-trade/coin-exchange-api/exchange-api/internal/handler/admin"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.CheckSignMiddleware, serverCtx.TokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/admin",
					Handler: admin.AdminAddHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/admin",
					Handler: admin.AdminUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/admin",
					Handler: admin.AdminDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/admin",
					Handler: admin.AdminSearchHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/admin/detail",
					Handler: admin.AdminDetailHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/admin"),
	)
}
