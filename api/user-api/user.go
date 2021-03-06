package main

import (
	"flag"
	"fmt"
	"frozen-go-project/api/user-api/internal/config"
	"frozen-go-project/api/user-api/internal/handler"
	"frozen-go-project/api/user-api/internal/middleware"
	"frozen-go-project/api/user-api/internal/svc"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")
var port = flag.Int("port", 0, "rpc port")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	if *port <= 0 {
		panic("should provide listen port")
	}
	c.Port = *port
	c.Log.Path = fmt.Sprintf("%s/%d", c.Log.Path, c.Port)

	ctx := svc.NewServiceContext(c)
	//跨域支持
	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(rest.CorsHandler()))
	defer server.Stop()
	server.Use(middleware.CheckSign(ctx))
	server.Use(middleware.CheckBan(ctx))

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
