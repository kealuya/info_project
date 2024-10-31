package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/wumansgy/goEncrypt/des"
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

	msg := "床前明月光，疑是地上霜，举头望明月，低头思故乡"
	desSecretKey := "szhtszht"
	base64Text, err := des.DesCbcEncryptHex([]byte(msg), []byte(desSecretKey), []byte("12345678"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("DES模式加密后的base64密文为:\n%s\n", base64Text)

	plaintext, err := des.DesCbcDecryptByHex(base64Text, []byte(desSecretKey), []byte("12345678"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("DES模式解密后:\n%s\n", string(plaintext))

}
