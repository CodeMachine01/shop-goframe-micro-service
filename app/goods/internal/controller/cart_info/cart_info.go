package cart_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service/app/goods/api/pbentity"
	"shop-goframe-micro-service/app/goods/internal/dao"
	"shop-goframe-micro-service/app/goods/internal/model/entity"
	"shop-goframe-micro-service/utility"
	"shop-goframe-micro-service/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCartInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCartInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.CartInfoGetListReq) (res *v1.CartInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.CartInfoListResponse{
		List:  make([]*pbentity.CartInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.CartInfo, consts.GetListFail)
	// 查询总数
	total, err := dao.CartInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	cartRecords, err := dao.CartInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		Where("user_id", req.UserId).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range cartRecords {
		var cart entity.CartInfo
		if err := record.Struct(&cart); err != nil {
			continue
		}

		var pbCart pbentity.CartInfo
		if err := gconv.Struct(cart, &pbCart); err != nil {
			continue
		}

		pbCart.CreatedAt = utility.SafeConvertTime(cart.CreatedAt)
		pbCart.UpdatedAt = utility.SafeConvertTime(cart.UpdatedAt)

		response.List = append(response.List, &pbCart)
	}

	return &v1.CartInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CartInfoCreateReq) (res *v1.CartInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CartInfo, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.CartInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.CartInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CartInfoDeleteReq) (res *v1.CartInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.CartInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.CartInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.CartInfoDeleteRes{}, nil // 返回空结构体
}
