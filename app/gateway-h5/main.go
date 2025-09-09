package main

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	_ "shop-goframe-micro-service/app/gateway-h5/internal/packed"
	"shop-goframe-micro-service/utility"

	"github.com/gogf/gf/v2/os/gctx"

	"shop-goframe-micro-service/app/gateway-h5/internal/cmd"
)

func main() {
	var ctx = gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}
	var address = conf.String()
	//使得gRPC客户端可以通过服务名发现服务地址
	grpcx.Resolver.Register(etcd.New(address))

	// 创建HTTP服务
	s := g.Server()

	//设置CORS头
	s.Use(utility.MiddlewareCORS)

	cmd.Main.Run(ctx)
}
