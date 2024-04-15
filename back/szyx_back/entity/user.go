package entity

//kdxf  用户信息
type UserInfo struct {
	Id         string `json:"id" description:"id" `
	UserId     string `json:"userId" description:"用户ID" `
	UserName   string `json:"userName" description:"用户姓名" `
	UserMobile string `json:"userMobile" description:"用户手机号" `
	Sex        string `json:"sex" description:"性别" `
	BankCard   string `json:"bankCard" description:"银行卡号" `
	BankName   string `json:"bankName" description:"银行名称" `
	CorpName   string `json:"corpName" description:"企业名称" `
	CorpCode   string `json:"corpCode" description:"企业code"`
	CreateTime string `json:"createTime" description:"创建时间"`
	Creater    string `json:"creater" description:"创建人"`
	Bz1        string `json:"bz1" description:"备注1"`
	Bz2        string `json:"bz2" description:"备注2"`
	Bz3        string `json:"bz3" description:"备注3"`
}

type LoginResponseData struct {
	UserInfo UserInfo `json:"userInfo"`
	Msg      string   `json:"msg"`
	Result   bool     `json:"result"`
}

//=====================================================================================

//登录入参信息
type LoginRequestPar struct {
	LoginName     string `json:"loginName" description:"用户名" `      // 用户名
	LoginPassword string `json:"loginPassword" description:"密码" `   // 密码
	Captcha       string `json:"captcha" description:"验证码" `        // 验证码
	CaptchaId     string `json:"captchaId" description:"验证码Id" `    // 验证码ID
	CorpCode      string `json:"companyId" description:"企业代码" `     //企业代码
	LoginWay      int    `json:"loginWay" description:"登录方式" `      // 1 -- 企业用户  2-- 手机号邮箱 3--二维码
	EquipmentType string `json:"equipmentType" description:"设备类型" ` //登录设备类型 PC,APP
}
