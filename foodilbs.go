package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/config"
	geoServer "github.com/foodi-org/foodi-lbs-server/internal/server/geo"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/dev.yaml", "use the dev.yaml config file")

func main() {
	var c config.Config
	flag.Parse()
	dir, _ := os.Getwd()

	if err := svc.NewServiceContext(c, dir, *configFile); err != nil {
		panic(err)
	}

	// 注册服务
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {

		foodi_lbs_server.RegisterGeoServer(grpcServer, geoServer.NewGeoServer(svc.Svc())) // lbs位置服务
		foodi_lbs_server.RegisterDeliveryServer(grpcServer)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting foodi-lbs rpc server at %s...\n", c.ListenOn)
	s.Start()
}
