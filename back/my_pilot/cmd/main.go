package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"strings"

	//_ "my_pilot/pkg/api_jd_iop"
	//"my_pilot/pkg/api_jd_iop/product"
	//_ "my_pilot/pkg/api_lxwl_shop"
	_ "my_pilot/pkg/api_stb_shop"
	"my_pilot/pkg/api_stb_shop/product"
)

func main() {
	stb()
	select {}
}

func stb() {
	fmt.Println("stb is ok")
	pageNum, err := product.GetPageNum()
	if err != nil {
		fmt.Println(err)
	}

	logs.Info("%+v", pageNum)

	for _, p := range pageNum.Result {
		skus, err := product.GetSkus(p.PageNum)
		if err != nil {
			fmt.Println(err)
		}

		for _, sku := range strings.Split(skus.Result, ",") {
			detail, err := product.GetDetail(sku)
			if err != nil {
				fmt.Println(err)
			}
			func(a any) {}(detail)
			price, err := product.GetPrice(skus.Result)
			if err != nil {
				fmt.Println(err)
			}

			logs.Info("%+v", price)

		}

	}

}

//	func lxwl() {
//		fmt.Println("lxwl is ok")
//		pools, err := lxwl_product.GetPools()
//		if err != nil {
//			fmt.Println(err)
//		}
//		logs.Info("%+v", pools)
//		for _, v := range pools.Result {
//			skus, err := lxwl_product.GetSkus(v.No)
//			if err != nil {
//				fmt.Println(err)
//			}
//			logs.Info("%+v", skus)
//			detail, err := lxwl_product.GetDetail(skus.Result)
//			if err != nil {
//				fmt.Println(err)
//			}
//			logs.Info("%+v", detail)
//		}
//	}
//func jd() {
//	r, err := product.GetPageNum()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(r)
//}
