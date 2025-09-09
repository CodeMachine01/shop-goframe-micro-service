package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 创建收货地址请求
type ConsigneeInfoCreateReq struct {
	g.Meta    `path:"/consignee" method:"post" tags:"收货地址管理" summary:"创建收获地址"`
	UserId    uint32 `json:"userId"    description:"用户id"`        //用户id
	IsDefault uint32 `json:"isDefault"  description:"默认地址1 非默认0"` // 默认地址1 非默认0
	Name      string `json:"name"            description:"姓名"`    // 姓名
	Phone     string `json:"phone"           description:"电话号"`   // 电话号
	Province  string `json:"province"    description:"省"`         // 省
	City      string `json:"city"         description:"市"`        // 市
	Town      string `json:"town"            description:"县区"`    // 县区
	Street    string `json:"street"        description:"街道乡镇"`    // 街道乡镇
	Detail    string `json:"detail"      description:"地址详情"`      // 地址详情
}

// 创建收货地址响应
type ConsigneeInfoCreateRes struct {
	Id uint32 `json:"id" dc:"收货地址ID"`
}

// 获取收获地址列表请求
type ConsigneeInfoGetListReq struct {
	g.Meta `path:"/consignee" method:"get" tags:"收货地址管理" summary:"获取收货地址列表"`
	Page   uint32 `json:"page" v:"min:1" dc:"页码" default:"1"`
	Size   uint32 `json:"size" v:"max:100" dc:"每页数量" default:"10"`
}

// 获取收货地址列表响应
type ConsigneeInfoGetListRes struct {
	List  []*ConsigneeInfoItem `json:"list" dc:"收获地址列表"`
	Page  uint32               `json:"page" dc:"收获地址列表"`
	Size  uint32               `json:"size" dc:"每页数量"`
	Total uint32               `json:"total" dc:"总数"`
}

// 收获地址项
type ConsigneeInfoItem struct {
	Id        uint32                 `json:"id"          description:"id"`        // id
	UserId    uint32                 `json:"userId"     description:"用户id"`       //用户id
	IsDefault uint32                 `json:"isDefault"  description:"默认地址1 非默认0"` // 默认地址1 非默认0
	Name      string                 `json:"name"      description:"姓名"`          // 姓名
	Phone     string                 `json:"phone"       description:"电话号"`       // 电话号
	Province  string                 `json:"province"     description:"省"`        // 省
	City      string                 `json:"city"            description:"市"`     // 市
	Town      string                 `json:"town"        description:"县区"`        // 县区
	Street    string                 `json:"street"        description:"街道乡镇"`    // 街道乡镇
	Detail    string                 `json:"detail"        description:"地址详情"`    // 地址详情
	CreatedAt *timestamppb.Timestamp `json:"createdAt" description:"创建时间"`        // 创建时间
	UpdatedAt *timestamppb.Timestamp `json:"updatedAt"  description:"更新时间"`       // 更新时间
	DeletedAt *timestamppb.Timestamp `json:"deletedAt"  description:"删除时间"`       // 删除时间
}

// 更新收货地址请求
type ConsigneeInfoUpdateReq struct {
	g.Meta    `path:"/consignee" method:"put" tags:"收货地址管理" summary:"更新收货地址"`
	Id        uint32 `json:"id"          description:"id"`        // id
	IsDefault uint32 `json:"isDefault"  description:"默认地址1 非默认0"` // 默认地址1 非默认0
	Name      string `json:"name"      description:"姓名"`          // 姓名
	Phone     string `json:"phone"       description:"电话号"`       // 电话号
	Province  string `json:"province"     description:"省"`        // 省
	City      string `json:"city"            description:"市"`     // 市
	Town      string `json:"town"        description:"县区"`        // 县区
	Street    string `json:"street"        description:"街道乡镇"`    // 街道乡镇
	Detail    string `json:"detail"        description:"地址详情"`    // 地址详情
}

// 更新收获地址响应
type ConsigneeInfoUpdateRes struct {
	Id uint32 `json:"id" dc:"收货地址ID"`
}

// 删除收获地址请求
type ConsigneeInfoDeleteReq struct {
	g.Meta `path:"/consignee" method:"delete" tags:"收货地址管理" summary:"删除收货地址"`
	Id     uint32 `json:"id" v:"required" description:"收货地址id"`
}

// 删除收货地址响应
type ConsigneeInfoDeleteRes struct {
}
