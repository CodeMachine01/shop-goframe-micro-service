package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
	"shop-goframe-micro-service/app/interaction/internal/controller/collection_info"
	"shop-goframe-micro-service/app/interaction/internal/controller/comment_info"
	"shop-goframe-micro-service/app/interaction/internal/controller/praise_info"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "interaction grpc server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(grpcx.Server.UnaryValidate)}...,
			)
			s := grpcx.Server.New(c)
			collection_info.Register(s)
			comment_info.Register(s)
			praise_info.Register(s)
			s.Run()
			return nil
		},
	}
)
