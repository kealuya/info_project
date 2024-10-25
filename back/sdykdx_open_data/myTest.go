package main

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"strings"
)

type TextProcessorInterface interface {
	Process(text string) string
}

// PlainText 是一种基本文本处理器
type PlainText struct{}

func (p *PlainText) Process(text string) string {
	return text
}

// TextDecorator 是装饰器的基础类型
//type TextDecorator struct {
//	TextProcessorInterface
//}

// NewlineDecorator 为文本增加换行符
type NewlineDecorator struct {
	TextProcessorInterface
}

func (d *NewlineDecorator) Process(text string) string {
	return d.TextProcessorInterface.Process(text) + "\n"
}

// UppercaseDecorator 将文本转换为大写
type UppercaseDecorator struct {
	TextProcessorInterface
}

func (d *UppercaseDecorator) Process(text string) string {
	processedText := d.TextProcessorInterface.Process(text)
	return strings.ToUpper(processedText)
}

func main() {
	// 构建基础 URL    https://sso.ccmu.edu.cn/serviceValidate?service=http://gzpt.ccmu.edu.cn/personCommon/portal_cas&ticket=xxxxxxxx
	baseURL := "https://sso.ccmu.edu.cn/serviceValidate"

	//u, _ := url.Parse(baseURL)
	//// 构建查询参数
	//queryParams := url.Values{}
	// 读取casurl
	casUrl, _ := config.String("casurl")
	//queryParams.Set("service", casUrl)
	//queryParams.Set("ticket", ticket)
	//u.RawQuery = queryParams.Encode()
	//logs.Info("请求地址get::" + u.String())

	baseURL = baseURL + "?service=" + casUrl + "&ticket=" + "xxxxxxxxx"
	logs.Info("请求地址get::" + baseURL)

	//redirect := "http://gzpt.ccmu.edu.cn"
	//params := url.Values{}
	//params.Add("isOk", "true")
	//params.Add("user", "User")
	//params.Add("name", "Name")
	//params.Add("employee_number", "EmployeeNumber")
	//u, _ := url.Parse(redirect)
	//
	//desSecretKey := "szhtszht"
	//hexText, err := des.DesCbcEncryptHex([]byte(params.Encode()), []byte(desSecretKey), nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//u.RawQuery = "rawQuery=" + hexText
	//logs.Info("redirect_meta::" + redirect + "?rawQuery=" + params.Encode())
	//logs.Info("redirect_dec::" + u.String())
	//
	//plain := &PlainText{}
	//textWithNewline := &NewlineDecorator{plain}
	//textWithDecorations := &UppercaseDecorator{textWithNewline}
	//
	//result := textWithDecorations.Process("Hello, Decorator!")
	//fmt.Print(result) // Outputs: "HELLO, DECORATOR!\n"

	//plaintext1, err := des.DesCbcDecryptByHex("1a36d81d29656ed2ca7100dffac5a679747aa0ccc42208f1b76b6667e5926b756c929f7f804dec411fb1abf85e4c55720a4d9d48237c1de7", []byte(desSecretKey), nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("DES模式解密后:\n%s\n", string(plaintext1))

}
