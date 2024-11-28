package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	_ "my_pilot/pkg/api_jd_iop"
	jd_product "my_pilot/pkg/api_jd_iop/product"
	_ "my_pilot/pkg/api_stb_shop"
	stb_message "my_pilot/pkg/api_stb_shop/message"
	stb_product "my_pilot/pkg/api_stb_shop/product"
	"strings"
)

type User struct {
	name string
}

func (u User) Name() {
	fmt.Println("Name:", u.name)
}

func (u *User) PointName() {
	fmt.Println("PointName:", u.name)
}

func printUser() {
	u := User{name: "user1"}

	defer u.Name()
	defer u.PointName()

	u.name = "user2"
}

func main() {
	//message()
	//jd()
	//select {}
	printUser()

}

func message() {

	getMessage, err := stb_message.GetMessage("0", "1,2,3,4,5")
	if err != nil {
		fmt.Println(err)
	}

	pushMessage := stb_message.PushMessage{
		PriceUpdate: func(result stb_message.PriceUpdateResult) error {
			fmt.Println(result)
			return nil
		},
	}
	for _, msg := range getMessage.Result {
		fmt.Println(msg)
		err := pushMessage.RunLogicWithGetMessageResult(msg)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func stb() {
	fmt.Println("stb is ok")
	pageNum, err := stb_product.GetPageNum()
	if err != nil {
		fmt.Println(err)
	}

	logs.Info("%+v", pageNum)

	for _, p := range pageNum.Result {
		skus, err := stb_product.GetSkus(p.PageNum)
		if err != nil {
			fmt.Println(err)
		}

		for _, sku := range strings.Split(skus.Result, ",") {
			detail, err := stb_product.GetDetail(sku)
			if err != nil {
				fmt.Println(err)
			}
			func(a any) {}(detail)
			price, err := stb_product.GetPrice(skus.Result)
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
func jd() {
	_, err := jd_product.GetPageNum()
	if err != nil {
		fmt.Println(err)
	}

	getPageNumResult, err := jd_product.GetPageNum()
	if err != nil {
		return
	}

	//fmt.Printf("%#v", getPageNumResult)

	for _, sk := range getPageNumResult.Result {
		detail, err := jd_product.QuerySkuByPage(sk.PageNum, 1, "0")
		if err != nil {
			return
		}
		fmt.Println(detail)

		//jd_product.GetDetail()

	}

}
