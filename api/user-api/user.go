package main

import (
	"flag"
	"fmt"
	"frozen-go-project/api/user-api/internal/config"
	"frozen-go-project/api/user-api/internal/handler"
	"frozen-go-project/api/user-api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	//server.Use(func(next http.HandlerFunc) http.HandlerFunc {
	//	return func(w http.ResponseWriter, r *http.Request) {
	//		logx.Info("request ... ")
	//		next(w, r)
	//		logx.Infof("response ...%v ", w)
	//	}
	//})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
