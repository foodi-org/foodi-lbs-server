package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/config"
	deliverServer "github.com/foodi-org/foodi-lbs-server/internal/server/delivery"
	geoServer "github.com/foodi-org/foodi-lbs-server/internal/server/geo"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "dev.yaml", "use the dev.yaml config file")

func main() {
	var c = config.ServConf()
	flag.Parse()
	dir, _ := os.Getwd()

	if err := svc.NewServiceContext(c, dir, *configFile); err != nil {
		panic(err)
	}

	// 注册服务
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {

		// lbs位置服务
		foodi_lbs_server.RegisterGeoServer(grpcServer, geoServer.NewGeoServer(svc.Svc()))

		// 配送位置服务，作为服务分层注册示例
		foodi_lbs_server.RegisterDeliveryServer(grpcServer, deliverServer.NewDeliveryServer(svc.Svc()))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting foodi-lbs rpc server at %s...\n", c.ListenOn)
	s.Start()
}
