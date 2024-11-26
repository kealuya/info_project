package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	_ "my_pilot/pkg/api_jd_iop"
	"my_pilot/pkg/api_jd_iop/product"
	_ "my_pilot/pkg/api_lxwl_shop"
	lxwl_product "my_pilot/pkg/api_lxwl_shop/product"
)

func main() {
	select {}
}
func lxwl() {
	fmt.Println("lxwl is ok")
	pools, err := lxwl_product.GetPools()
	if err != nil {
		fmt.Println(err)
	}
	logs.Info("%+v", pools)
	for _, v := range pools.Result {
		skus, err := lxwl_product.GetSkus(v.No)
		if err != nil {
			fmt.Println(err)
		}
		logs.Info("%+v", skus)
		detail, err := lxwl_product.GetDetail(skus.Result)
		if err != nil {
			fmt.Println(err)
		}
		logs.Info("%+v", detail)
	}
}
func jd() {
	r, err := product.GetPageNum()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
