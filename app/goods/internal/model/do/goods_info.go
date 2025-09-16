// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsInfo is the golang structure of table goods_info for DAO operations like Where/Data.
type GoodsInfo struct {
	g.Meta           `orm:"table:goods_info, do:true"`
	Id               any         //
	Name             any         // 商品名字
	Images           any         // 支持单图,多图
	PicUrl           any         // 主图URL
	Price            any         // 价格(分)
	Level1CategoryId any         // 1级分类id
	Level2CategoryId any         // 2级分类id
	Level3CategoryId any         // 3级分类id
	Brand            any         // 品牌
	Stock            any         // 库存
	Sale             any         // 销量
	Tags             any         // 标签
	DetailInfo       any         // 商品详情
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	DeletedAt        *gtime.Time //
}
