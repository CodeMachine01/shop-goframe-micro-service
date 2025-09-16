package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_info "shop-goframe-micro-service/app/user/api/user_info/v1"
	"shop-goframe-micro-service/utility/middleware"

	"shop-goframe-micro-service/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.UserInfoReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		panic("用户ID类型错误或不存在")
	}
	grpcReq.Id = userId

	grpcRes, err := c.UserInfoClient.GetUserInfo(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 使用gconv转换响应
	res = &v1.UserInfoRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}

	return res, nil
}
