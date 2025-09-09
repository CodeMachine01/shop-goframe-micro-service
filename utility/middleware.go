package utility

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"google.golang.org/grpc"
	"time"
)

// 拦截器
func GrpcClientTimeout(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	//用带超时的上下文ctx，调用真实rpc
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

// CORS跨域资源共享 中间件
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,OPTIONS")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	//处理预检请求
	if r.Method == "OPTIONS" {
		r.Response.WriteHeader(204) //No Content
		return
	}
	r.Middleware.Next()
}
