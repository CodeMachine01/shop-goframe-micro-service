package consignee_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service/app/user/api/consignee_info/v1"
	"shop-goframe-micro-service/app/user/api/pbentity"
	"shop-goframe-micro-service/app/user/internal/consts"
	"shop-goframe-micro-service/app/user/internal/dao"
	"shop-goframe-micro-service/app/user/internal/model/entity"
	"shop-goframe-micro-service/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedConsigneeInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterConsigneeInfoServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.ConsigneeInfoCreateReq) (res *v1.ConsigneeInfoCreateRes, err error) {
	////定义一个实体对象，用于接收转换后的请求数据
	//var consigneeInfo *entity.ConsigneeInfo
	////将请求参数req转换为实体对象consigneeInfo
	//if err := gconv.Struct(req, &consigneeInfo); err != nil {
	//	return nil, err
	//}
	//错误类型 定制需要的样式
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.CreateFail)
	//向数据库中插入数据并获取自动生成的ID
	result, err := dao.ConsigneeInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	//返回创建成功响应，包含新创建的ID
	return &v1.ConsigneeInfoCreateRes{
		Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.ConsigneeInfoDeleteReq) (res *v1.ConsigneeInfoDeleteRes, err error) {
	//错误类型 定制需要的样式
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.DeleteFail)
	//根据ID从数据库中删除对应信息
	_, err = dao.ConsigneeInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	//返回删除成功的空响应
	return &v1.ConsigneeInfoDeleteRes{}, err
}

func (*Controller) Update(ctx context.Context, req *v1.ConsigneeInfoUpdateReq) (res *v1.ConsigneeInfoUpdateRes, err error) {
	////定义一个实体对象，用于接收转换后的请求数据
	//var consigneeInfo entity.ConsigneeInfo
	////将请求参数req转换为实体对象consigneeInfo
	//if err := gconv.Struct(req, &consigneeInfo); err != nil {
	//	return nil, err
	//}
	//错误类型 定制需要的样式
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.UpdateFail)

	//根据ID更新数据库中的信息
	_, err = dao.ConsigneeInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	//返回更新成功响应，包含被更新ID
	return &v1.ConsigneeInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.ConsigneeInfoGetListReq) (res *v1.ConsigneeInfoGetListRes, err error) {
	response := &v1.ConsigneeInfoListResponse{
		List:  make([]*pbentity.ConsigneeInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	//错误类型 定制需要的样式
	infoError := consts.InfoError(consts.ConsigneeInfo, consts.GetListFail)

	//查询总数
	total, err := dao.ConsigneeInfo.Ctx(ctx).Count() //gf的ORM框架的数据库查询 SELECT COUNT(*) as total FROM `consignee_info`;
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	//查询当前页数据
	consigneeRecords, err := dao.ConsigneeInfo.Ctx(ctx). //SELECT * FROM `consignee_info` LIMIT {req.Size} OFFSET {(req.Page - 1) * req.Size};
								Where("user_id", req.UserId).
								Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	//数据转换
	//在循环中替换手动赋值
	for _, record := range consigneeRecords {
		//数据库记录gdb.Record->实体结构体entity.ConsigneeInfo
		var consignee entity.ConsigneeInfo
		if err := record.Struct(&consignee); err != nil {
			continue
		}

		//实体结构体entity.ConsigneeInfo->Protobuf结构体pbentity.ConsigneeInfo
		var pbConsignee pbentity.ConsigneeInfo
		if err := gconv.Struct(consignee, &pbConsignee); err != nil {
			continue
		}

		//单独处理时间片段（因为gconv无法自动转换）
		pbConsignee.CreatedAt = utility.SafeConvertTime(consignee.CreatedAt)
		pbConsignee.UpdatedAt = utility.SafeConvertTime(consignee.UpdatedAt)
		pbConsignee.DeletedAt = utility.SafeConvertTime(consignee.DeletedAt)

		//添加到响应列表
		response.List = append(response.List, &pbConsignee)
	}

	return &v1.ConsigneeInfoGetListRes{Data: response}, nil
}
