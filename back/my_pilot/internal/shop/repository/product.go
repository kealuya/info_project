package repository

import (
	"github.com/gohouse/t"
	"my_pilot/common"
	jd_product "my_pilot/pkg/api_jd_vop/product"
)

func StoreProductDetail() {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)
	defer session.Commit()

	categories := make([]Category, 0)
	errWhere := session.Where("cat_class = ?", 2).Find(&categories)
	common.ErrorHandler(errWhere)

	for _, category := range categories {
		_, err := jd_product.QuerySkuByPage(t.New(category.CatId).String(), 1)
		common.ErrorHandler(err)
		//skuids := page.Result.SkuIds
		//pageCount := page.Result.PageCount
	}

	return
}
