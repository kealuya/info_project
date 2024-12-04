package main

import (
	_ "my_pilot/config"
	"my_pilot/internal/shop/handler"
	_ "my_pilot/pkg/api_jd_vop"
)

func main() {
	handler.StoreProductDetail()

}

func productDetailHandler() {
	//detail, err := jd_product.QuerySkuByPage(sk.PageNum, 1)
	//if err != nil {
	//	logs.Error(err)
	//}
}
