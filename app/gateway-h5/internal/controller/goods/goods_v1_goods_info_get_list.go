package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_info "shop-goframe-micro-service/app/goods/api/goods_info/v1"

	"shop-goframe-micro-service/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoGetList(ctx context.Context, req *v1.GoodsInfoGetListReq) (res *v1.GoodsInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &goods_info.GoodsInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.GoodsInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.GoodsInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	// 批量转换列表项
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
