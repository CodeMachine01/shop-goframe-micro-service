package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 商品详情
type GoodsInfoGetDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品管理" sm:"商品详情"`
	Id     uint32 `json:"id" v:"required" dc:"商品ID"`
}

type GoodsInfoGetDetailRes struct {
	*GoodsInfoItem // 复用列表项结构
}

// 商品分页查询
type GoodsInfoGetListReq struct {
	g.Meta `path:"/goods" method:"get" tags:"商品管理" sm:"商品分页列表"`
	Page   uint32 `json:"page" d:"1"  v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
	IsHot  uint32 `json:"is_hot" d:"0"  dc:"热门推荐(1 开启)"`
}

type GoodsInfoGetListRes struct {
	List  []*GoodsInfoItem `json:"list" dc:"商品列表"`
	Page  uint32           `json:"page" dc:"当前页码"`
	Size  uint32           `json:"size" dc:"每页数量"`
	Total uint32           `json:"total" dc:"总数"`
}

type GoodsInfoItem struct {
	Id               uint32                 `json:"id" dc:"商品ID"`
	Name             string                 `json:"name" dc:"商品名称"`
	PicUrl           string                 `json:"pic_url" dc:"主图"`
	Images           string                 `json:"images" dc:"支持单图,多图"`
	Price            uint64                 `json:"price" dc:"价格"`
	Sort             uint32                 `json:"sort" dc:"排序 倒序"`
	Level1CategoryId uint32                 `json:"level1_category_id" dc:"一级分类ID"`
	Level2CategoryId uint32                 `json:"level2_category_id" dc:"二级分类ID"`
	Level3CategoryId uint32                 `json:"level3_category_id" dc:"三级分类ID"`
	Brand            string                 `json:"brand" dc:"品牌"`
	Stock            uint32                 `json:"stock" dc:"库存"`
	Sale             uint32                 `json:"sale" dc:"销量"`
	Tags             string                 `json:"tags" dc:"标签"`
	DetailInfo       string                 `json:"detail_info" dc:"详情"`
	CreatedAt        *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt        *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
	DeletedAt        *timestamppb.Timestamp `json:"deleted_at" dc:"删除时间"`
}
