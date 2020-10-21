// Code generated by goctl. DO NOT EDIT!
// Source: user_rpc.proto

package main

import (
	"flag"
	"fmt"

	"frozen-go-project/rpc/user-rpc/internal/config"
	"frozen-go-project/rpc/user-rpc/internal/server"
	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user-rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	userRpcSrv := server.NewUserRpcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user_rpc.RegisterUserRpcServer(grpcServer, userRpcSrv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
