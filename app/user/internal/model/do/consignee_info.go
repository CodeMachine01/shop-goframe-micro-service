// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsigneeInfo is the golang structure of table consignee_info for DAO operations like Where/Data.
type ConsigneeInfo struct {
	g.Meta    `orm:"table:consignee_info, do:true"`
	Id        any         // 收货地址表
	UserId    any         //
	IsDefault any         // 默认地址1 非默认0
	Name      any         // 姓名
	Phone     any         // 电话号
	Province  any         // 省
	City      any         // 市
	Town      any         // 县区
	Street    any         // 街道乡镇
	Detail    any         // 地址详情
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
