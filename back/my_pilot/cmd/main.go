package main

import (
	"fmt"
	_ "my_pilot/pkg/api_jd_iop"
	"my_pilot/pkg/api_jd_iop/product"
)

func main() {

	r, err := product.GetPageNum()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	select {}
}
