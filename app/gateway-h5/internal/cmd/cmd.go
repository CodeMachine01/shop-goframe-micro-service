package cmd

import (
	"context"
	"shop-goframe-micro-service/app/gateway-h5/internal/controller/banner"
	"shop-goframe-micro-service/app/gateway-h5/internal/controller/goods"
	"shop-goframe-micro-service/app/gateway-h5/internal/controller/interaction"
	"shop-goframe-micro-service/app/gateway-h5/internal/controller/order"
	"shop-goframe-micro-service/app/gateway-h5/internal/controller/user"
	"shop-goframe-micro-service/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway-h5 server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 创建控制器实例一次，重复使用
			userController := user.NewV1()
			goodsController := goods.NewV1()
			bannerController := banner.NewV1()
			interactionController := interaction.NewV1()
			orderController := order.NewV1()
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(
						userController.UserInfoLogin,
						userController.UserInfoRegister,
						goodsController.CategoryInfoGetAll,
						goodsController.CategoryInfoGetList,
						goodsController.GoodsInfoGetDetail,
						goodsController.GoodsInfoGetList,
						bannerController,
					)
				})
				// 需要JWT验证的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTAuth)
					group.Bind(
						//需要认证的接口
						userController.ConsigneeInfoCreate,
						userController.ConsigneeInfoDelete,
						userController.ConsigneeInfoGetList,
						userController.ConsigneeInfoUpdate,
						userController.UserInfo,
						userController.UserInfoUpdatePassword,

						goodsController.CartInfoGetList,
						goodsController.CartInfoCreate,
						goodsController.CartInfoDelete,
						interactionController,
						orderController,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
