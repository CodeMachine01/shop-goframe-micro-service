package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service/app/gateway-h5/api/goods/v1"
	cart_info "shop-goframe-micro-service/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service/utility/middleware"
)

func (c *ControllerV1) CartInfoCreate(ctx context.Context, req *v1.CartInfoCreateReq) (res *v1.CartInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &cart_info.CartInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		panic("用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId
	// 调用gRPC服务
	grpcRes, err := c.CartInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CartInfoCreateRes{Id: grpcRes.Id}, nil
}
