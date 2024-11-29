package main

import (
	"encoding/json"
	"fmt"
	_ "my_pilot/pkg/api_jd_vop"
	jd_product "my_pilot/pkg/api_jd_vop/product"
)

func main() {
	//message()
	jdVop()
	select {}

}
func jdVop() {

	getPageNumResult, err := jd_product.GetPageNum()
	if err != nil {
		return
	}

	b, err := json.Marshal(getPageNumResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", b)

	for _, sk := range getPageNumResult.Result {
		detail, err := jd_product.QuerySkuByPage(sk.PageNum, 2)
		if err != nil {
			return
		}
		fmt.Printf("%v:: %+v\n", sk.PageNum, detail)
	}

}

//100022117855
//func message() {
//
//	getMessage, err := stb_message.GetMessage("0", "1,2,3,4,5")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	pushMessage := stb_message.PushMessage{
//		PriceUpdate: func(result stb_message.PriceUpdateResult) error {
//			fmt.Println(result)
//			return nil
//		},
//	}
//	for _, msg := range getMessage.Result {
//		fmt.Println(msg)
//		err := pushMessage.RunLogicWithGetMessageResult(msg)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//}

//func stb() {
//	fmt.Println("stb is ok")
//	pageNum, err := stb_product.GetPageNum()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	logs.Info("%+v", pageNum)
//
//	for _, p := range pageNum.Result {
//		skus, err := stb_product.GetSkus(p.PageNum)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		for _, sku := range strings.Split(skus.Result, ",") {
//			detail, err := stb_product.GetDetail(sku)
//			if err != nil {
//				fmt.Println(err)
//			}
//			func(a any) {}(detail)
//			price, err := stb_product.GetPrice(skus.Result)
//			if err != nil {
//				fmt.Println(err)
//			}
//
//			logs.Info("%+v", price)
//
//		}
//
//	}
//
//}
