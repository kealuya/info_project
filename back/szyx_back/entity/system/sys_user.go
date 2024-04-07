package system

import "szyx_back/entity"

type SysUser struct {
	UUID        string `description:"用户Id" json:"uuid"`          // 用户Code
	Username    string `description:"用户登录名称" json:"userName"`    //
	Password    string `description:"用户登录密码" json:"password" `   // 用户登录密码
	OldPassword string `description:"旧密码" json:"oldPassword"`    //旧密码
	NickName    string `description:"用户昵称" json:"nickName"`      // 用户昵称
	SideMode    string `description:"用户侧边主题" json:"sideMode"`    // 用户侧边主题
	HeaderImg   string `description:"用户头像" json:"headerImg"`     // 用户头像
	BaseColor   string `description:"基础颜色" json:"baseColor"`     // 基础颜色
	ActiveColor string `json:"activeColor"`                      // 活跃颜色
	AuthorityId string `description:"用户权限Id" json:"authorityId"` // 用户角色ID
	AuthName    string `description:"权限名称" json:"authName"`
	//Authority      SysAuthority   `json:"authority"`
	//Authorities    []SysAuthority `json:"authorities"`
	CorpCode       string `description:"公司编码" json:"corpCode"`
	CorpName       string `description:"公司名称" json:"corpName"`
	UserRole       string `json:"userRole"`
	IsManage       string `json:"isManage"`
	ClearPassword  string `json:"clearPassword"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	PersonType     string `description:"人员类型" json:"personType"`
	DirectLeader   string `description:"直属领导" json:"directLeader"`
	Sex            string `description:"性别" json:"sex"`
	EntryTime      string `description:"入职时间" json:"entryTime"`
	DeptName       string `description:"部门名称" json:"deptName"`
	JobName        string `description:"职位名称" json:"jobName"`
	JobTitleId     string `json:"jobTitleId"`
	PostName       string `description:"职务名称" json:"postName"`
	Level          string `description:"级别" json:"level"`
	EmpCode        string `description:"工号" json:"empCode"`
	OpeningTime    string `description:"企业开户时间" json:"openingTime"`
	ValidDate      string `description:"有效时间" json:"validDate"`
	DeptId         string `json:"deptId"`
	ChineseName    string `json:"chineseName"`
	IdCard         string `json:"idCard"`
	OpenState      string `json:"openState"`
	PostbandId     string `json:"postbandId"`
	Desc           string `json:"desc"`
	BZ             string `json:"bz"`
	EnglishName    string `json:"englishName"`
	AppDeviceToken string `json:"appDeviceToken"`
}

type SysUserAuthority struct {
	SysEmpCode     string `description:"用户Id" json:"sysEmpCode"`
	SysAuthorityId string `description:"权限Id" json:"sysAuthorityId"`
	CorpCode       string `json:"corpCode"`
}

type SetUserAuthorities struct {
	EmpCode      string   `description:"用户Id" json:"empCode"`
	AuthorityIds []string `description:"角色Id" json:"authorityIds"`
	CorpCode     string   `json:"corpCode"`
}

type LoginRequestPar struct {
	LoginName     string `description:"用户名" json:"loginName"`      // 用户名
	LoginPassword string `description:"密码" json:"loginPassword"`   // 密码
	Captcha       string `description:"验证码" json:"captcha"`        // 验证码
	CaptchaId     string `description:"验证码Id" json:"captchaId"`    // 验证码ID
	CorpCode      string `description:"企业代码" json:"companyId"`     //企业代码
	LoginWay      int    `description:"登录方式" json:"loginWay"`      // 1 -- 企业用户  2-- 手机号邮箱 3--二维码
	EquipmentType string `description:"设备类型" json:"equipmentType"` //登录设备类型 PC,APP
}

type GetMessageCaptcha struct {
	Phone string `json:"phone"`
}

type MessageAlibaba struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
type MessageAlibabaReturn struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type MessageLogin struct {
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	ValidateCode  string `json:"validateCode"`
	EquipmentType string `description:"设备类型" json:"equipmentType"` //登录设备类型 PC,APP
}

type ModifyPhoneNum struct {
	NewPhone      string `json:"newPhone"`
	OldPhone      string `json:"oldPhone"`
	ValidateCode  string `json:"validateCode"`
	EquipmentType string `description:"设备类型" json:"equipmentType"` //登录设备类型 PC,APP
}

type PhoneVerification struct {
	Phone         string `json:"phone"`
	ValidateCode  string `json:"validateCode"`
	EquipmentType string `description:"设备类型" json:"equipmentType"` //登录设备类型 PC,APP
}

type MessageLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type PageUserList struct {
	Page     int    `description:"页码" json:"page"`
	PageSize int    `description:"每页显示条数" json:"pageSize"`
	KeyWord  string `description:"搜索关键字" json:"keyWord"`
	CorpCode string `description:"企业编码" json:"corpCode"`
}

type DeleteUser struct {
	EmpCode  string `description:"用户Id" json:"empCode"`
	CorpCode string `json:"corpCode"`
}

type GetUserInfo struct {
	EmpCode  string `description:"用户Id" json:"empCode"`
	CorpCode string `json:"corpCode"`
}

type JwtDecrypt struct {
	JWT      string `description:"jwt密文" json:"jwt"` // jwt密文
	CorpCode string `description:"企业编码" json:"corpCode"`
}

type ResultInfo struct {
	JWT          string               `description:"jwt" json:"jwt"`                   // access_token
	RefreshToken string               `description:"refreshToken" json:"refreshToken"` //refreshToken
	User         entity.LoginUserInfo `description:"用户信息" json:"user"`
}
type FunctionConfig struct {
	Id               string `description:"自增id" json:"id"`
	YwTypeID         string `json:"ywTypeID"`
	YwType           string `json:"ywType"`
	YwTypeName       string `json:"ywTypeName"`
	YwTypeInfo       string `json:"ywTypeInfo"`
	ControllerTypeBz string `json:"controllerTypeBz"`
	ConfigValue      string `json:"configValue"`
	ConfigExplain    string `json:"configExplain"`
	Creater          string `description:"creater" json:"creater"`       //type:string   comment:                                         version:2022-07-11 10:57
	CreateTime       string `description:"createTime" json:"createTime"` //type:string   comment:                                         version:2022-07-11 10:57
	CorpCode         string `description:"corpCode" json:"corpCode"`     //type:string   comment:
	DeptId           string `description:"deptId" json:"deptId"`         //type:string   comment:
	DeptName         string `description:"DeptName" json:"deptName"`     //type:string   comment:
}

type UserConfig struct {
	Id          string `description:"自增id" json:"id"`
	EmpCode     string `description:"用户ID" json:"empCode"`
	EmpName     string `description:"用户名" json:"empName"`
	TravelPower string `description:"差旅权限" json:"travelPower"`
	OtherPower  string `description:"其他权限" json:"otherPower"`
	Creater     string `description:"" json:"creater"`
	CreateTime  string `description:"" json:"createTime"`
	CorpCode    string `description:"企业编码" json:"corpCode"`
	Bz          string `description:"" json:"bz"`
}

type UserConfigReq struct {
	UserInfo    []UserInfo `description:"用户信息" json:"userInfo"`
	TravelPower string     `description:"差旅权限" json:"travelPower"`
	OtherPower  string     `description:"其他权限" json:"otherPower"`
	Creater     string     `json:"creater"`
	CreateTime  string     `json:"createTime"`
	Bz          string     `json:"bz"`
}

type UserListCount struct {
	Total    int       `description:"总数" json:"total"`
	UserList []SysUser `description:"用户列表" json:"list"`
}

type UploadUser struct {
	CorpCode string `json:"corpCode"`
	File     string `json:"file"`
}

type InsertUserInfo struct {
	CorpCode     string   `json:"corpCode"`
	EmpCode      string   `json:"empCode"`
	EmpName      string   `json:"empName"`
	Password     string   `json:"password"`
	Mobile       string   `json:"mobile"`
	AuthorityId  []string `description:"角色Id" json:"authorityId"`
	Creater      string   `json:"creater"`
	CreateTime   string   `json:"createTime"`
	IdCard       string   `json:"idCard"`
	DeptId       string   `json:"deptId"`
	DeptName     string   `json:"deptName"`
	JobName      string   `json:"jobName"`
	JobId        string   `json:"jobId"`
	PostBandName string   `json:"postBandName"`
	PostBandId   string   `json:"postBandId"`
	Modifier     string   `json:"modifier"`
	ModifyTime   string   `json:"modifyTime"`
}

type UserInfo struct {
	ChineseName string `json:"chineseName"`
	EmpCode     string `json:"empCode"`
	Phone       string `json:"phone"`
	DeptId      string `json:"deptId"`
	DeptName    string `json:"deptName"`
	JobTitleId  string `json:"jobTitleId"`
	IdCard      string `json:"idCard"`
	PostBandId  string `json:"postBandId"`
	PostBand    string `json:"postBand"`
	CorpCode    string `json:"corpCode"`
	JobTitle    string `json:"jobTitle"`
	Department  string `json:"department"`
}

type UserSignAccount struct {
	Address     string `json:"address"`
	DeptName    string `json:"deptName"`
	Email       string `json:"email"`
	IdentNo     string `json:"identNo"`
	MobilePhone string `json:"mobilePhone"`
	PersonName  string `json:"personName"`
	EmpCode     string `json:"empCode"`
	CorpCode    string `json:"corpCode"`
}

type CfcaSignResponse struct {
	ErrorMessage string `json:"errorMessage"`
	Head         struct {
		RetCode    string `json:"retCode"`
		RetMessage string `json:"retMessage"`
		TxTime     string `json:"txTime"`
	} `json:"head"`
	NotSendPwd int64 `json:"notSendPwd"`
	Person     struct {
		Address            string `json:"address"`
		AnXinSignEmail     string `json:"anXinSignEmail"`
		AuthenticationMode string `json:"authenticationMode"`
		Email              string `json:"email"`
		IdentNo            string `json:"identNo"`
		IdentTypeCode      string `json:"identTypeCode"`
		IsOpenSM2          int64  `json:"isOpenSM2"`
		MobilePhone        string `json:"mobilePhone"`
		PersonName         string `json:"personName"`
		UsedEmailLogin     int64  `json:"usedEmailLogin"`
		UsedMobileLogin    int64  `json:"usedMobileLogin"`
		UserID             string `json:"userId"`
		EmpCode            string `json:"empCode"`
		CorpCode           string `json:"corpCode"`
	} `json:"person"`
}

type CfcaSignInfo struct {
	EmpCode            string `json:"empCode"`
	CfcaId             string `json:"cfcaId"`
	PersonName         string `json:"personName"`
	IdCardNo           string `json:"idCardNo"`
	MobilePhone        string `json:"mobilePhone"`
	OpenState          string `json:"openState"`
	Email              string `json:"email"`
	AuthenticationMode string `json:"authenticationMode"`
	Address            string `json:"address"`
	CorpCode           string `json:"corpCode"`
}

type PersonIDCardResp struct {
	Msg    string `json:"msg"`
	Result bool   `json:"result"`
}

type EmpJobRank struct {
	CorpCode   string `description:"企业id"  json:"corpCode"`
	EmpCode    string `description:"员工工号" json:"empCode"`
	PostBandId string `description:"职级id" json:"postBandId"`
	PostBand   string `description:"职级名称" json:"postBand"`
}
type EmpResponseData struct {
	Msg    string `json:"msg"`
	Result bool   `json:"result"`
}
