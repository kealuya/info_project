package product

import (
	"fmt"
	"my_pilot/pkg/api_jd_iop"
	"strings"
)

// GetSellPrice 查询商品售卖价
func GetSellPrice(sku []string, skuInfos []struct{ int64 int }, queryExts string) (*api_jd_iop.Response[[]SellPriceResult], error) {

	config := api_jd_iop.GetJdIopConfig()
	url := config["jd_iop"]["base_url"] + "price/getSellPrice"
	token := config["token"]["access_token"]
	client := api_jd_iop.GlobalClient.R()

	// 构造请求参数
	// queryExts 为英文半角分隔的多个枚举值，枚举值不同，本接口的出参不同。
	//枚举值如下：
	//price //大客户默认价格(根据合同类型查询价格)。
	//marketPrice //市场价。
	//containsTax //税率。
	//出参增加tax,taxPrice,nakedPrice 3个字段。
	//nakedPrice//未税价(nakedPrice这个扩展参数废弃了，废弃了，废弃了)。代表本客户为未税价下单模式（需要运营在主数据提前配置）；
	//出参增加nakedPrice字段 (加此入参时，出参price也是未税价，此时price= nakedPrice)
	var skuInfosString string
	if skuInfos != nil {
		marshal, err := api_jd_iop.JsonEncoder.Marshal(skuInfos)
		if err != nil {
			return nil, err
		}
		skuInfosString = string(marshal)
	}
	formData := map[string]string{
		"token":     token,
		"sku":       strings.Join(sku, ","),
		"skuInfos":  skuInfosString,
		"queryExts": queryExts,
	}

	resultResp := &api_jd_iop.Response[[]SellPriceResult]{}
	resp, errClient := client.SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(resultResp).
		ForceContentType("application/json").
		Post(url)
	if errClient != nil {
		return nil, errClient
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode())
	}
	return resultResp, nil
}

// SellPriceResult 价格结果结构体
type SellPriceResult struct {
	// skuId
	SkuID int64 `json:"skuId"`

	// 京东价。仅供参考
	JDPrice float64 `json:"jdPrice"`

	// 京东销售价
	// 含税价下单模式时返回含税单价
	// 未税价下单模式时返回未税单价
	Price float64 `json:"price"`

	// 价格描述
	PriceDesc string `json:"priceDesc"`

	// 京东的前台划线价
	// 入参中的queryExts中包含marketPrice时，输出此字段
	// 现在只有图书频道能露出，其他的因政策原因已不允许展示
	// 仅供参考
	MarketPrice float64 `json:"marketPrice"`

	// 税率
	// 当queryExts中包含containsTax时，出参中有此字段
	// 例如：此值为16时，代表税率为"16%"
	Tax float64 `json:"tax"`

	// 预估税额
	// 当queryExts中包含containsTax时，出参中有此字段
	TaxPrice float64 `json:"taxPrice"`

	// 未税价
	// 当queryExts中包含containsTax或nakedPrice时，出参中有此字段
	// ① 入参containsTax：此字段代表含税价订单体系下，单品的预估未税价，
	//    仅作页面展示，不作为订单、票面中最终的未税单价，
	//    因为下单后会有运费分摊、发票尾差校验等处理逻辑
	// ② 入参nakedPrice：此字段代表未税价订单体系下，单品的未税单价，
	//    此时price= nakedPrice
	NakedPrice float64 `json:"nakedPrice"`
}
