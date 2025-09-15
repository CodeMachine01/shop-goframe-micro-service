package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	_ "shop-goframe-micro-service/app/goods/internal/packed"
	"shop-goframe-micro-service/app/goods/utility/goodsRedis"

	"github.com/gogf/gf/v2/os/gctx"

	"shop-goframe-micro-service/app/goods/internal/cmd"
)

func main() {
	var ctx = gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}
	// 初始化Redis
	if err := goodsRedis.InitGoodsRedis(ctx); err != nil {
		g.Log().Fatal(ctx, "Redis初始化失败:", err)
		os.Exit(1)
	}
	var address = conf.String()
	grpcx.Resolver.Register(etcd.New(address))
	cmd.Main.Run(ctx)
}
