package entity

//用户信息
type UserInfo struct {
	Corp     string `json:"Corp"`
	Username string `json:"username"`
	EmpCode  string `json:"EmpCode"`
	Tel      string `json:"tel"`
	IdCard   string `json:"idCard"`
	Dept     string `json:"dept"`
}

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
