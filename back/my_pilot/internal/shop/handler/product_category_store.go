package handler

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gohouse/t"
	"github.com/jinzhu/copier"
	"log"
	"my_pilot/internal/shop/repository"
	jd_product "my_pilot/pkg/api_jd_vop/product"
)

var Set = make(map[string]any)

func getCategoryMe(categories *[]repository.Category, pageNum string) {
	requestCategory, err := jd_product.GetCategory(pageNum)
	if err != nil {
		logs.Error(err)
		log.Panicln(err)
	}
	fmt.Printf("%#v\n", requestCategory.Result)
	category := repository.Category{Status: 1}
	err = copier.Copy(&category, requestCategory.Result)
	if err != nil {
		logs.Error(err)
		log.Panicln(err)
	}
	*categories = append(*categories, category)
	Set[t.New(category.CatId).String()] = struct{}{}
	_, ok := Set[t.New(category.ParentId).String()]
	if category.CatClass != 0 && !ok {
		Set[t.New(category.ParentId).String()] = struct{}{}
		getCategoryMe(categories, t.New(category.ParentId).String())
	}
}

func StoreProductCategory() {

	categories := make([]repository.Category, 0, 100)

	// 获取商品池
	getPageNumResult, err := jd_product.GetPageNum()
	if err != nil {
		logs.Error(err)
	}
	for _, sk := range getPageNumResult.Result {

		getCategoryMe(&categories, sk.PageNum)

		if len(categories) >= 100 {
			repository.InsertCategory(categories)
			categories = categories[:0]
		}
	}
	if len(categories) > 0 {
		repository.InsertCategory(categories)
		categories = categories[:0]
	}

}
