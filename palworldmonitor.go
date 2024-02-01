package main

import (
	"flag"
	"fmt"

	"github.com/KiClover/palworld-status-rpc-monitor/internal/config"
	"github.com/KiClover/palworld-status-rpc-monitor/internal/server"
	"github.com/KiClover/palworld-status-rpc-monitor/internal/svc"
	"github.com/KiClover/palworld-status-rpc-monitor/types/palworldmonitor"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/palworldmonitor.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		palworldmonitor.RegisterPalworldmonitorServer(grpcServer, server.NewPalworldmonitorServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
