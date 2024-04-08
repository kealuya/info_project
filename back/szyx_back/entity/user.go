package entity

//kdxf  用户信息
type UserInfo struct {
	Id         string `description:"id" json:"id"`
	UserId     string `description:"用户ID" json:"userId"`
	UserName   string `description:"用户姓名" json:"userName"`
	UserMobile string `description:"用户手机号" json:"userMobile"`
	Sex        string `description:"性别" json:"sex"`
	BankCard   string `description:"银行卡号" json:"bankCard"`
	BankName   string `description:"银行名称" json:"bankName"`
	Base
}

type LoginResponseData struct {
	UserInfo UserInfo `json:"userInfo"`
	Msg      string   `json:"msg"`
	Result   bool     `json:"result"`
}

//=====================================================================================

//登录入参信息
type LoginRequestPar struct {
	LoginName     string `description:"用户名" json:"loginName"`      // 用户名
	LoginPassword string `description:"密码" json:"loginPassword"`   // 密码
	Captcha       string `description:"验证码" json:"captcha"`        // 验证码
	CaptchaId     string `description:"验证码Id" json:"captchaId"`    // 验证码ID
	CorpCode      string `description:"企业代码" json:"companyId"`     //企业代码
	LoginWay      int    `description:"登录方式" json:"loginWay"`      // 1 -- 企业用户  2-- 手机号邮箱 3--二维码
	EquipmentType string `description:"设备类型" json:"equipmentType"` //登录设备类型 PC,APP
}
