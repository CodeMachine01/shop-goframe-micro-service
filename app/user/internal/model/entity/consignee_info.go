// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsigneeInfo is the golang structure for table consignee_info.
type ConsigneeInfo struct {
	Id        int         `json:"id"        orm:"id"         description:"收货地址表"`      // 收货地址表
	UserId    int         `json:"userId"    orm:"user_id"    description:""`           //
	IsDefault int         `json:"isDefault" orm:"is_default" description:"默认地址1 非默认0"` // 默认地址1 非默认0
	Name      string      `json:"name"      orm:"name"       description:"姓名"`         // 姓名
	Phone     string      `json:"phone"     orm:"phone"      description:"电话号"`        // 电话号
	Province  string      `json:"province"  orm:"province"   description:"省"`          // 省
	City      string      `json:"city"      orm:"city"       description:"市"`          // 市
	Town      string      `json:"town"      orm:"town"       description:"县区"`         // 县区
	Street    string      `json:"street"    orm:"street"     description:"街道乡镇"`       // 街道乡镇
	Detail    string      `json:"detail"    orm:"detail"     description:"地址详情"`       // 地址详情
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`       // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`       // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`       // 删除时间
}
