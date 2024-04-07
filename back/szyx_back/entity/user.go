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

//登录成功返回值
type LoginUserInfo struct {
	UserName     string   `description:"用户名" json:"userName"`
	EmpCode      string   `description:"用户Id" json:"empCode"`
	Password     string   `description:"密码" json:"password"`
	NickName     string   `description:"昵称" json:"nickName"`
	AuthorityIds []string `description:"权限Id组" json:"authority"`
	Phone        string   `description:"电话" json:"phone"`
	CorpCode     string   `description:"企业编码" json:"corpCode"`
	UserRole     string   `description:"用户角色" json:"userRole"`
	UserRoleName []string `description:"用户角色名称" json:"userRoleName"`
	IsManage     string   `description:"是否是管理员" json:"isManage"`
	IdCard       string   `description:"身份证号" json:"idCard"`
	DeptName     string   `json:"deptName"`
	DeptId       string   `json:"deptId"`
	PostBandId   string   `description:"职级ID" json:"postBandId"`
	PostBand     string   `description:"职级" json:"postBand"`
	JobTitle     string   `description:"职位" json:"jobTitle"`
	JobTitleId   string   `description:"职位Id" json:"jobTitleId"`
	OpenState    string   `description:"开户状态" json:"openState"`
	AppKey       string   `json:"appKey"`
	AppSceret    string   `json:"appSceret"`
	CustomerNo   string   `json:"customerNo"`
	Email        string   `json:"email"`
	JobName      string   `json:"jobName"`
	BZ           string   `json:"bz"`
	CorpName     string   `json:"corpName"`
}

type LoginResponseData struct {
	UserInfo LoginUserInfo `json:"userInfo"`
	Msg      string        `json:"msg"`
	Result   bool          `json:"result"`
}
