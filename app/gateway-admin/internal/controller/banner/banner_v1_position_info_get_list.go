package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	position_info "shop-goframe-micro-service/app/banner/api/position_info/v1"

	"shop-goframe-micro-service/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) PositionInfoGetList(ctx context.Context, req *v1.PositionInfoGetListReq) (res *v1.PositionInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &position_info.PositionInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.PositionInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.PositionInfoGetListRes{
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
