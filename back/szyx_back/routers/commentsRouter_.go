package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "DeleteInvoiceToWb",
			Router:           "/deleteInvoiceToWb",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "GetFpInputTypeList",
			Router:           "/getFpInputTypeList",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "GetInvOcrResultByPdfWb",
			Router:           "/getInvOcrResultByPdfWb",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "GetInvOcrResultWb",
			Router:           "/getInvOcrResultWb",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "GetInvoiceImgUrl",
			Router:           "/getInvoiceImgUrl",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "GetInvoiceWbList",
			Router:           "/getInvoiceWbList",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "InvVerifyWb",
			Router:           "/invVerifyWb",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "UpdateInvoiceStates",
			Router:           "/updateInvoiceStates",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"] = append(beego.GlobalControllerRouter["szyx_back/controllers:InvoiceToSzyxApi"],
		beego.ControllerComments{
			Method:           "UploadBase64",
			Router:           "/uploadBase64",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:JwtController"] = append(beego.GlobalControllerRouter["szyx_back/controllers:JwtController"],
		beego.ControllerComments{
			Method:           "GetToken",
			Router:           "/getToken",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:JwtController"] = append(beego.GlobalControllerRouter["szyx_back/controllers:JwtController"],
		beego.ControllerComments{
			Method:           "RefreshToken",
			Router:           "/refreshToken",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:SysLoginController"] = append(beego.GlobalControllerRouter["szyx_back/controllers:SysLoginController"],
		beego.ControllerComments{
			Method:           "CreateCaptcha",
			Router:           "/createCaptcha",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:SysLoginController"] = append(beego.GlobalControllerRouter["szyx_back/controllers:SysLoginController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["szyx_back/controllers:SysLoginController"] = append(beego.GlobalControllerRouter["szyx_back/controllers:SysLoginController"],
		beego.ControllerComments{
			Method:           "SetUserInfo",
			Router:           "/setUserInfo",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
