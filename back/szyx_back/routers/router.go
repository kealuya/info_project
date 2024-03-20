// @APIVersion 1.0.0
// @Title 数字营销后端接口
// @Description 数字营销 Api
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"szyx_back/controllers"
)

func init() {

	namespace :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/log",
				beego.NSInclude(&controllers.LogInfoController{}),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(&controllers.SysLoginController{}),
			),
			beego.NSNamespace("/token",
				beego.NSInclude(&controllers.JwtController{}),
			),
			//----------发票识别、验真相关- 开始 ----------------
			beego.NSNamespace("/invoice",
				beego.NSRouter("/getInvoiceWbList", &controllers.InvoiceToSzyxApi{}, "post:GetInvoiceWbList"),
			),
			beego.NSNamespace("/invoice",
				beego.NSRouter("/deleteInvoiceToWb", &controllers.InvoiceToSzyxApi{}, "post:DeleteInvoiceToWb"),
			),
			beego.NSNamespace("/invoice",
				beego.NSRouter("/getInvoiceImgUrl", &controllers.InvoiceToSzyxApi{}, "post:GetInvoiceImgUrl"),
			),
			beego.NSNamespace("/invoice",
				beego.NSRouter("/uploadBase64", &controllers.InvoiceToSzyxApi{}, "post:UploadBase64"),
			),
			beego.NSNamespace("/invoice",
				beego.NSRouter("/getFpInputTypeList", &controllers.InvoiceToSzyxApi{}, "post:GetFpInputTypeList"),
			),
			beego.NSNamespace("/invoice",
				beego.NSRouter("/invVerifyWb", &controllers.InvoiceToSzyxApi{}, "post:InvVerifyWb"),
			),
			beego.NSNamespace("/invoice",
				beego.NSRouter("/getInvOcrResultWb", &controllers.InvoiceToSzyxApi{}, "post:GetInvOcrResultWb"),
			),
			beego.NSNamespace("/invoice",
				beego.NSRouter("/getInvOcrResultByPdfWb", &controllers.InvoiceToSzyxApi{}, "post:GetInvOcrResultByPdfWb"),
			),
			//----------发票识别、验真相关 - 结束 ----------------
		)
	//注册 namespace
	beego.AddNamespace(namespace)
	beego.SetStaticPath("/swagger", "swagger")
}
