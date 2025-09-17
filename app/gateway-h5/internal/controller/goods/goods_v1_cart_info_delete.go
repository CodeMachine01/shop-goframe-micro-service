package goods

import (
	"context"
	cart_info "shop-goframe-micro-service/app/goods/api/cart_info/v1"

	"shop-goframe-micro-service/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) CartInfoDelete(ctx context.Context, req *v1.CartInfoDeleteReq) (res *v1.CartInfoDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.CartInfoClient.Delete(ctx, &cart_info.CartInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.CartInfoDeleteRes{}, nil
}
