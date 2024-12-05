package handler

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gohouse/t"
	"github.com/jinzhu/copier"
	"my_pilot/internal/shop/repository"
	jd_product "my_pilot/pkg/api_jd_vop/product"
	"sync"
)

func StoreProductDetail() {
	categories := repository.GetCategoryByCatClass(2)

	dbHandler := make(chan jd_product.DetailResult, 200)
	categoryHandler := make(chan repository.Category, 10)

	for i := 0; i < 10; i++ {
		go func() {
			handleCategoryPoolGetProduct(categoryHandler, dbHandler)
		}()
	}
	wgDb := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wgDb.Add(1)
		go func() {
			defer wgDb.Done()
			handlerDbInsertProduct(dbHandler)
		}()
	}

	for _, category := range categories {
		categoryHandler <- category
	}
	close(categoryHandler)

	wgDb.Wait()
	close(dbHandler)

	return
}

func handlerDbInsertProduct(dbHandler chan jd_product.DetailResult) {
	var products []repository.Product
	for detailResult := range dbHandler {
		dr := repository.Product{}
		copier.Copy(&dr, detailResult)
		products = append(products, dr)
		if len(products) == 100 {
			repository.InsertProducts(products)
			products = products[:0]
		}
	}
	repository.InsertProducts(products)
}

func handleCategoryPoolGetProduct(categoryHandler chan repository.Category, dbHandler chan jd_product.DetailResult) {
	for category := range categoryHandler {
		c, err := jd_product.QuerySkuByPage(t.New(category.CatId).String(), 1)
		if err != nil {
			logs.Error(err)
			continue
		}
		if !c.Success {
			logs.Error(fmt.Errorf("业务请求不正确:%#v:%#v", category, c))
			continue
		}
		pageCount := c.Result.PageCount
		for i := 1; i <= pageCount; i++ {
			categoryPool, err := jd_product.QuerySkuByPage(t.New(category.CatId).String(), i)
			if err != nil {
				logs.Error(err)
				continue
			}
			if !c.Success {
				logs.Error(fmt.Errorf("业务请求不正确:%#v:%#v", category, c))
				continue
			}
			skuids := categoryPool.Result.SkuIds
			for _, sku := range skuids {
				detail, err := jd_product.GetDetail(sku, "")
				if err != nil {
					logs.Error("GetDetail请求不正确", err)
					continue
				}
				if !detail.Success {
					logs.Error("GetDetail业务请求不正确::%#v", detail)
					continue
				}
				var d jd_product.DetailResult = detail.Result
				dbHandler <- d
			}
		}
	}
}
