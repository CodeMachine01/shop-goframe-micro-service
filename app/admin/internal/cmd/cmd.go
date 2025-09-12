package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
	"shop-goframe-micro-service/app/admin/internal/controller/admin_info"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "admin login grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			//向配置中加入自定义选项
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary( // 将多个一元拦截器组合成一个链式拦截器
					grpcx.Server.UnaryValidate), //请求参数验证拦截器，用于自动校验grpc请求的输入参数是否符合Protobuf中定义规则
			}...) //Unary为一元，Stream为流式，一元是指单个请求到单个调用
			s := grpcx.Server.New(c) //创建grpc实例
			admin_info.Register(s)   //将admin_info注册到grpc服务器
			s.Run()                  //阻塞当前goroutine，启动grpc服务器并开始监听地址
			return nil
		},
	}
)
