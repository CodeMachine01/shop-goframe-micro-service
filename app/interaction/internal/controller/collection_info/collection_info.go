package collection_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service/app/interaction/api/collection_info/v1"
	"shop-goframe-micro-service/app/interaction/api/pbentity"
	"shop-goframe-micro-service/app/interaction/internal/consts"
	"shop-goframe-micro-service/app/interaction/internal/dao"
	"shop-goframe-micro-service/app/interaction/internal/model/entity"
	"shop-goframe-micro-service/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCollectionInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCollectionInfoServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CollectionInfoCreateReq) (res *v1.CollectionInfoCreateRes, err error) {
	//错误类型 定制需要的样式
	infoError := consts.InfoError(consts.CollectionInfo, consts.CreateFail)
	//向数据库中插入数据并获取自动生成的ID
	result, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	//返回创建成功响应，包含新创建的ID
	return &v1.CollectionInfoCreateRes{
		Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error) {
	//错误类型 定制需要的样式
	infoError := consts.InfoError(consts.CollectionInfo, consts.DeleteFail)
	//根据ID从数据库中删除对应信息
	_, err = dao.CollectionInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	//返回删除成功的空响应
	return &v1.CollectionInfoDeleteRes{}, err
}

func (*Controller) GetList(ctx context.Context, req *v1.CollectionInfoGetListReq) (res *v1.CollectionInfoGetListRes, err error) {
	response := &v1.CollectionInfoListResponse{
		List:  make([]*pbentity.CollectionInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	//错误类型 定制需要的样式
	infoError := consts.InfoError(consts.CollectionInfo, consts.GetListFail)

	//查询总数
	total, err := dao.CollectionInfo.Ctx(ctx).Count() //gf的ORM框架的数据库查询 SELECT COUNT(*) as total FROM `consignee_info`;
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	//查询当前页数据
	consigneeRecords, err := dao.CollectionInfo.Ctx(ctx). //SELECT * FROM `consignee_info` LIMIT {req.Size} OFFSET {(req.Page - 1) * req.Size};
								Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		//记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	//数据转换
	//在循环中替换手动赋值
	for _, record := range consigneeRecords {
		//数据库记录gdb.Record->实体结构体entity.CollectionInfo
		var consignee entity.CollectionInfo
		if err := record.Struct(&consignee); err != nil {
			continue
		}

		//实体结构体entity.CollectionInfo->Protobuf结构体pbentity.CollectionInfo
		var pbConsignee pbentity.CollectionInfo
		if err := gconv.Struct(consignee, &pbConsignee); err != nil {
			continue
		}

		//单独处理时间片段（因为gconv无法自动转换）
		pbConsignee.CreatedAt = utility.SafeConvertTime(consignee.CreatedAt)
		pbConsignee.UpdatedAt = utility.SafeConvertTime(consignee.UpdatedAt)

		//添加到响应列表
		response.List = append(response.List, &pbConsignee)
	}

	return &v1.CollectionInfoGetListRes{Data: response}, nil
}
