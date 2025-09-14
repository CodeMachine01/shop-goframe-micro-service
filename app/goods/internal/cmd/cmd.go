package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
	"shop-goframe-micro-service/app/goods/internal/controller/goods_images"
	"shop-goframe-micro-service/app/goods/internal/controller/goods_info"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "goods grpc server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(grpcx.Server.UnaryValidate)}...,
			)
			s := grpcx.Server.New(c)
			goods_info.Register(s)
			goods_images.Register(s)
			s.Run()
			return nil
		},
	}
)
