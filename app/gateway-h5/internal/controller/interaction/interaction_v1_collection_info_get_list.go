package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	collection "shop-goframe-micro-service/app/interaction/api/collection_info/v1"

	"shop-goframe-micro-service/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CollectionInfoGetList(ctx context.Context, req *v1.CollectionInfoGetListReq) (res *v1.CollectionInfoGetListRes, err error) {
	//使用gconv自动转换结构体
	grpcReq := &collection.CollectionInfoGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	//调用gRPC服务
	grpcRes, err := c.CollectionInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	//转换响应
	res = &v1.CollectionInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	//批量转换列表项
	if err := gconv.Struct(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
