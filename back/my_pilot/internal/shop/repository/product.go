package repository

import (
	"my_pilot/common"
	"time"
)

type Product struct {
	ID           int64     `xorm:"pk autoincr bigint(20) 'id' comment('主键ID')" json:"id"`
	Sku          int64     `xorm:"bigint(50) notnull unique 'sku' comment('商品编号')" json:"sku"`
	Name         string    `xorm:"varchar(255) notnull 'name' comment('商品名称')" json:"name"`
	BrandName    string    `xorm:"varchar(100) notnull index 'brand_name' comment('品牌名称')" json:"brandName"`
	Category     string    `xorm:"varchar(100) notnull 'category' comment('分类')" json:"category"`
	State        int       `xorm:"tinyint(1) notnull default 0 index 'state' comment('主站上下架状态(1上架 0下架)')" json:"state"`
	Status       int       `xorm:"tinyint(1) notnull default 0 index 'status' comment('商品状态(0下架 1上架)')" json:"status"`
	ImagePath    string    `xorm:"varchar(255) notnull 'image_path' comment('主图')" json:"imagePath"`
	WareQD       string    `xorm:"text notnull 'ware_qd' comment('包装清单')" json:"wareQD"`
	Introduction string    `xorm:"text notnull 'introduction' comment('商品详情页大字段')" json:"introduction"`
	SaleUnit     string    `xorm:"varchar(50) null 'sale_unit' comment('售卖单位')" json:"saleUnit,omitempty"`
	Weight       string    `xorm:"varchar(50) null 'weight' comment('重量')" json:"weight,omitempty"`
	ProductArea  string    `xorm:"varchar(100) null 'product_area' comment('产地')" json:"productArea,omitempty"`
	Param        string    `xorm:"text null 'param' comment('规格参数')" json:"param,omitempty"`
	UPC          string    `xorm:"varchar(50) null 'upc' comment('UPC码')" json:"upc,omitempty"`
	CreatedAt    time.Time `xorm:"created 'created_at' comment('创建时间')" json:"createdAt"`
	UpdatedAt    time.Time `xorm:"updated 'updated_at' comment('更新时间')" json:"updatedAt"`
	IsDeleted    bool      `xorm:"tinyint(1) notnull default 0 'is_deleted' comment('是否删除(0未删除 1已删除)')" json:"isDeleted"`
}

// TableName 指定表名
func (p *Product) TableName() string {
	return "product"
}

func InsertProducts(products []Product) {
	session := dbEngine.NewSession()
	defer session.Close()

	_, err := session.Insert(&products)
	if err != nil {
		session.Rollback()
		common.ErrorHandler(err)
		return
	}

	errSessionCommit := session.Commit()
	common.ErrorHandler(errSessionCommit)
	return
}
