package entity

//验证码struct
type Captcha struct {
	CaptchaId string `json:"captchaId"` //id
	PicPath   string `json:"picPath"`   //图片地址
}
