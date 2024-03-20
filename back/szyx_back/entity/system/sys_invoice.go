package system

type InvoiceWbList struct {
	EmpCode     string `json:"empCode"`
	CompanyCode string `description:"企业编码" json:"companyCode"`
	Row         string `json:"row"`
	Page        string `json:"page"`
	Uuid        string `json:"uuid"`
	CorpCode    string `json:"corpCode"`
	Type        string `description:"发票类型" json:"type"`
	Zt          string `description:"状态，0未使用，1已使用" json:"zt"`
	Dzfph       string `json:"dzfph"` //发票唯一标识
}

type InvoiceStates struct {
	Zt     string   `json:"zt"` //状态值  0未使用 1已使用
	Ywbh   string   `json:"ywbh"`
	Dzfphs []string `json:"dzfphs"` //电子发票号
	QyId   string   `json:"qyId"`
}
type ReturnOrderList struct {
	Data   ReturnData `json:"data"`
	Result string     `json:"result"`
	Msg    string     `json:"msg"`
}
type ReturnData struct {
	List []OrderLists `json:"list"`
}
type OrderLists struct {
	ShowDate    string      `json:"showDate"`
	OrderList   []OrderInfo `json:"orderList"`
	OrderXLList string      `json:"orderXLList"`
}
type OrderInfo struct {
	Applyno               string      `json:"applyno"`
	ApplyReason           string      `json:"applyReason"`
	EmpCode               string      `json:"empCode"`
	Orderno               string      `json:"orderno"`
	PlanBeginDate         string      `json:"planBeginDate"`
	PlanEndDate           string      `json:"planEndDate"`
	SpendTime             string      `json:"spendTime"`
	FromStation           string      `json:"fromStation"`
	ToStation             string      `json:"toStation"`
	TrainNo               string      `json:"trainNo"`
	OrderStatus           string      `json:"orderStatus"`
	PaymentType           string      `json:"paymentType"`
	PassengerName         string      `json:"passengerName"`
	CertificateCode       string      `json:"certificateCode"`
	SeatClass             string      `json:"seatClass"`
	SeatNo                string      `json:"seatNo"`
	GmtCreate             string      `json:"gmtCreate"`
	TotalAmount           string      `json:"totalAmount"`
	TicketStatus          string      `json:"ticketStatus"`
	PassengerTicketNo     string      `json:"passengerTicketNo"`
	TicketPriceCost       string      `json:"ticketPriceCost"`
	TicketPrice           string      `json:"ticketPrice"`
	ItemId                string      `json:"itemId"`
	IsUsed                string      `json:"isUsed"`
	OriginItemId          string      `json:"originItemId"`
	TrainChange           string      `json:"trainChange"`
	FlightChangeList      interface{} `json:"flightChangeList"`
	FlightInsuranceInfo   interface{} `json:"flightInsuranceInfo"`
	TrainRefund           string      `json:"trainRefund"`
	FlightRefund          string      `json:"flightRefund"`
	TrainYdInsuranceInfo  string      `json:"trainYdInsuranceInfo"`
	TrainGqdInsuranceInfo string      `json:"trainGqdInsuranceInfo"`
	Type                  string      `json:"type"`
	IsPersonal            string      `json:"isPersonal"`
	Originorderno         string      `json:"originorderno"`
	CityName              string      `json:"cityName"`
	HotelName             string      `json:"hotelName"`
	HotelAddress          string      `json:"hotelAddress"`
	GmtOccupancyTime      string      `json:"gmtOccupancyTime"`
	GmtLeaveTime          string      `json:"gmtLeaveTime"`
	OccupancyNights       string      `json:"occupancyNights"`
	RoomName              string      `json:"roomName"`
	RoomPriceDetail       string      `json:"roomPriceDetail"`
	RoomPrice             string      `json:"roomPrice"`
	BedTypeName           string      `json:"bedTypeName"`
	StartPort             string      `json:"startPort"`
	EndPort               string      `json:"endPort"`
	StartCityName         string      `json:"startCityName"`
	EndCityName           string      `json:"endCityName"`
	AirlineCompany        string      `json:"airlineCompany"`
	AirlineCompanyCode    string      `json:"airlineCompanyCode"`
	FlightNo              string      `json:"flightNo"`
	FlightDiscount        string      `json:"flightDiscount"`
	FlightAmount          string      `json:"flightAmount"`
	ServiceCharge         string      `json:"serviceCharge"`
	CapitalCost           string      `json:"capitalCost"`
	FuelCost              string      `json:"fuelCost"`
	PlusPrice             string      `json:"plusPrice"`
	FlightTotalAmount     string      `json:"flightTotalAmount"`
	IsViolation           string      `json:"isViolation"`
	ViolationContent      string      `json:"violationContent"`
	ViolationReason       string      `json:"violationReason"`
	ViolationReasonRemark string      `json:"violationReasonRemark"`
	OrderItemId           string      `json:"orderItemId"`
	NewOrderItemId        string      `json:"newOrderItemId"`
	ChangeFee             string      `json:"changeFee"`
	UpgradeCost           string      `json:"upgradeCost"`
	ChangeServiceCharge   string      `json:"changeServiceCharge"`
	ChangeAmount          string      `json:"changeAmount"`
	ChangePlusPrice       string      `json:"changePlusPrice"`
	RefundFinishDate      string      `json:"refundFinishDate"`
	RefundFee             string      `json:"refundFee"`
	RefundPlusPrice       string      `json:"refundPlusPrice"`
	RefundServiceCharge   string      `json:"refundServiceCharge"`
	RefundAmount          string      `json:"refundAmount"`
	OriginalOrderItemId   string      `json:"originalOrderItemId"`
	TicketAmount          string      `json:"ticketAmount"`
	TicketGmtCreate       string      `json:"ticketGmtCreate"`
	TicketGmtModified     string      `json:"ticketGmtModified"`
	TicketFee             string      `json:"ticketFee"`
	TaxFee                string      `json:"taxFee"`
	ServiceFee            string      `json:"serviceFee"`
	ItemStatus            string      `json:"itemStatus"`
	SegmentIndex          string      `json:"segmentIndex"`
	SegmentList           string      `json:"segmentList"`
	Destination           string      `json:"destination"`
	TicketNo              string      `json:"ticketNo"`
	Fromtypecode          string      `json:"fromtypecode"`
	Gqorderno             string      `json:"gqorderno"`
	Tporderno             string      `json:"tporderno"`
}

//请求发票列表出参
type ReturnInvoiceList struct {
	Data ReturnInvoiceData `json:"data"`
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
}
type ReturnInvoiceData struct {
	Count string         `json:"count"`
	Data  []InvoiceLists `json:"data"`
}
type InvoiceLists struct {
	Bsqrbh       string `json:"bsqrbh"`
	Bsqrxm       string `json:"bsqrxm"`
	Cc           string `json:"cc"`
	Cfdd         string `json:"cfdd"`
	Cfsj         string `json:"cfsj"`
	City         string `json:"city"`
	Cjrq         string `json:"cjrq"`
	Dddd         string `json:"dddd"`
	Dfdw         string `json:"dfdw"`
	Dfyh         string `json:"dfyh"`
	Dfzh         string `json:"dfzh"`
	Dzfph        string `json:"dzfph"`
	Fjid         string `json:"fjid"`
	Flag         string `json:"flag"`
	Fpdm         string `json:"fpdm"`
	Fphm         string `json:"fphm"`
	Fpygbh       string `json:"fpygbh"`
	Fwlb         string `json:"fwlb"`
	Gnfl         string `json:"gnfl"`
	InvName      string `json:"invName"`
	InvSource    string `json:"invSource"`
	IsFpYx       string `json:"isFpYx"`
	Issh         string `json:"issh"`
	IsverifyInv  bool   `json:"isverifyInv"`
	Isyz         string `json:"isyz"`
	Jjflkmbh     string `json:"jjflkmbh"`
	Kprq         string `json:"kprq"`
	Message      string `json:"message"`
	Pro          string `json:"pro"`
	Shrbh        string `json:"shrbh"`
	Shrxm        string `json:"shrxm"`
	Sqrbh        string `json:"sqrbh"`
	Sqrxm        string `json:"sqrxm"`
	Sqshrbh      string `json:"sqshrbh"`
	Sqshrxm      string `json:"sqshrxm"`
	Type         string `json:"type"`
	UUID         string `json:"uuid"`
	VerifyStatus string `json:"verifyStatus"`
	Xm           string `json:"xm"`
	Ygbh         string `json:"ygbh"`
	Ywbh         string `json:"ywbh"`
	Ywlx         string `json:"ywlx"`
	Ywlxmc       string `json:"ywlxmc"`
	Yydh         string `json:"yydh"`
	Yzm          string `json:"yzm"`
	Yzzt         string `json:"yzzt"`
	Zje          string `json:"zje"`
	ZjeTax       string `json:"zjeTax"`
	Zt           string `json:"zt"`
	Filename     string `json:"filename"`
}

/**
网报发票识别入参,调用网报的入参
*/
type GetTokenWbResp struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	Words string `json:"words"`
	Data  struct {
		EmpCode  string `json:"empCode"`
		CorpCode string `json:"corpCode"`
		Token    string `json:"token"`
	}
}

type ReturnInvoiceByIdList struct {
	Code string `json:"code"`
	Data []struct {
		Cc            string `json:"cc"`
		Cfdd          string `json:"cfdd"`
		Cfsj          string `json:"cfsj"`
		City          string `json:"city"`
		Cjrq          string `json:"cjrq"`
		Consumetypeid string `json:"consumetypeid"`
		Cph           string `json:"cph"`
		Dddd          string `json:"dddd"`
		Ddsj          string `json:"ddsj"`
		Dfdw          string `json:"dfdw"`
		Dfyh          string `json:"dfyh"`
		Dfzh          string `json:"dfzh"`
		Dzfph         string `json:"dzfph"`
		Fjid          string `json:"fjid"`
		Fpdm          string `json:"fpdm"`
		Fphm          string `json:"fphm"`
		Fplx          string `json:"fplx"`
		Fply          string `json:"fply"`
		Fwlb          string `json:"fwlb"`
		Gnfl          string `json:"gnfl"`
		InvNames      []struct {
			Dzfph string `json:"dzfph"`
			Fwlb  string `json:"fwlb"`
			Xmmc  string `json:"xmmc"`
			Zje   string `json:"zje"`
		} `json:"invNames"`
		InvSource  string `json:"inv_source"`
		Isclfdzfp  string `json:"isclfdzfp"`
		Isyz       string `json:"isyz"`
		JjflkmBH   string `json:"jjflkmBH"`
		Kpfsbh     string `json:"kpfsbh"`
		Filepath   string `json:"filepath"`
		Kpnr       string `json:"kpnr"`
		Kprq       string `json:"kprq"`
		Lrr        string `json:"lrr"`
		Pb         string `json:"pb"`
		Pickey2    string `json:"pickey2"`
		Pickey3    string `json:"pickey3"`
		Pro        string `json:"pro"`
		Scsj       string `json:"scsj"`
		Spfwmc     string `json:"spfwmc"`
		TxFilename string `json:"tx_filename"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Xcsj       string `json:"xcsj"`
		Xm         string `json:"xm"`
		Ygbh       string `json:"ygbh"`
		Ywbh       string `json:"ywbh"`
		Yzm        string `json:"yzm"`
		Zje        string `json:"zje"`
		ZjeTax     string `json:"zje_tax"`
		Zjh        string `json:"zjh"`
		Zt         string `json:"zt"`
	} `json:"data"`
	Msg   string      `json:"msg"`
	Words interface{} `json:"words"`
}

type PDFTojpgObs struct {
	PDFUrl string `json:"pdfUrl"`
}

type PDFTojpgObsReq struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data string `json:"data"`
}

//网报上传图片入参（手动录入发票）
type UploadBase64Request struct {
	QyId    string `json:"qyId"`    //企业ID
	EmpCode string `json:"empCode"` //用户编号
	ImgUrl  string `json:"imgUrl"`  //图片华为云obs地址
}

//网报获取手动输入发票类型列表入参
type GetFpInputTypeListRequest struct {
	QyId    string `json:"qyId"`    //企业ID
	EmpCode string `json:"empCode"` //用户编号
}

/**
网报保存发票入参
*/
type InvVerifyRequest struct {
	Fpdm          string `json:"fpdm"` //发票代码
	Fphm          string `json:"fphm"` //发票号码
	Kprq          string `json:"kprq"` //开票日期
	Jym           string `json:"jym"`  //校验码后六位
	Yzm           string `json:"yzm"`
	Pickey2       string `json:"pickey2"`
	Pickey3       string `json:"pickey3"`
	Isfpyzm       string `json:"isfpyzm"`
	Zje           string `json:"zje"` //金额
	ZjeTax        string `json:"zje_tax"`
	Fplx          string `json:"fplx"`
	FuncClassify  string `json:"funcClassify"`
	InvType       string `json:"invType"` //类型
	TravellerName string `json:"travellerName"`
	DepartTime    string `json:"departTime"`
	ArriveTime    string `json:"arriveTime"`
	DepartSite    string `json:"departSite"`
	ArriveSite    string `json:"arriveSite"`
	VehicleNum    string `json:"vehicleNum"`
	Pb            string `json:"pb"`
	Kpnr          string `json:"kpnr"`
	Kpdwmc        string `json:"kpdwmc"`
	Fwlb          string `json:"fwlb"`
	FileBase64    string `json:"fileBase64"` //图片的base64转码
	FileName      string `json:"fileName"`
	Ywlx          string `json:"ywlx"`
	InvSource     string `json:"invSource"`
	Fpfjuuid      string `json:"fpfjuuid"`
	Cph           string `json:"cph"`
	Zjh           string `json:"zjh"`
	Scsj          string `json:"scsj"`
	Xcsj          string `json:"xcsj"`
	QyId          string `json:"qyId"`        //企业ID
	EmpCode       string `json:"empCode"`     //企业ID
	PdfFilePath   string `json:"pdfFilePath"` //pdf文件识别地址
	FilePath      string `json:"filePath"`    //文件地址
}

/**
网报保存发票出参
*/
type InvVerifyResponse struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Words interface{} `json:"words"`
	Data  struct {
		Msg  interface{} `json:"msg"`
		Dzfp struct {
			Ztdm        interface{} `json:"ztdm"`
			Dwdm        interface{} `json:"dwdm"`
			Ischeck     string      `json:"ischeck"`
			Swlx        interface{} `json:"swlx"`
			Fphm        string      `json:"fphm"`
			Fkdzdh      string      `json:"fkdzdh"`
			Kpdwmc      string      `json:"kpdwmc"`
			Dj          int         `json:"dj"`
			Kpfsbh      string      `json:"kpfsbh"`
			Se          float64     `json:"se"`
			Fkdwmc      string      `json:"fkdwmc"`
			Fkfsbh      string      `json:"fkfsbh"`
			Ywbh        string      `json:"ywbh"`
			Kpdzdh      string      `json:"kpdzdh"`
			Dzfph       string      `json:"dzfph"`
			Je          float64     `json:"je"`
			Zje         int         `json:"zje"`
			Yzm         string      `json:"yzm"`
			Fpdm        string      `json:"fpdm"`
			Hyfl        string      `json:"hyfl"`
			Lrrq        string      `json:"lrrq"`
			Kprq        string      `json:"kprq"`
			Sl          int         `json:"sl"`
			Fkkhh       string      `json:"fkkhh"`
			Fpnr        string      `json:"fpnr"`
			Kpkhh       string      `json:"kpkhh"`
			Lrr         string      `json:"lrr"`
			Jbr         string      `json:"jbr"`
			Ssdq        interface{} `json:"ssdq"`
			Qyid        string      `json:"qyid"`
			Bz          interface{} `json:"bz"`
			Yydh        interface{} `json:"yydh"`
			Mongofjname interface{} `json:"mongofjname"`
			Pznm        interface{} `json:"pznm"`
			Lybh        interface{} `json:"lybh"`
			Kjnd        interface{} `json:"kjnd"`
			Pjlx        interface{} `json:"pjlx"`
			Bmbh        interface{} `json:"bmbh"`
			Xmbh        interface{} `json:"xmbh"`
			Kjqj        interface{} `json:"kjqj"`
			Zdr         interface{} `json:"zdr"`
			Isbx        interface{} `json:"isbx"`
		} `json:"dzfp"`
		Success bool `json:"success"`
	} `json:"data"`
}

/**
网报发票识别入参
*/
type InvOcrRequest struct {
	Base64     string `json:"base64"`   //图片的base64转码
	IsBase64   string `json:"isBase64"` //固定值，传1
	OcrInvType string `json:"ocrInvType"`
	QyId       string `json:"qyId"`     //企业编号
	EmpCode    string `json:"empCode"`  //用户ID
	FilePath   string `json:"filePath"` // 图片obs地址路径
}

/**
网报发票识别入参,调用网报的入参
*/
type InvOcrRequest_Wb struct {
	Base64     string `json:"base64"`   //图片的base64转码
	IsBase64   string `json:"isBase64"` //固定值，传1
	OcrInvType string `json:"ocrInvType"`
	QyId       string `json:"qyId"` //企业编号
}

/**
网报pdf,ofd发票识别入参
*/
type InvOcrByPdfRequest struct {
	Base64              string `json:"base64"`              //图片的base64转码
	IsBase64            bool   `json:"isBase64"`            //固定值，传false
	IsPdf               bool   `json:"isPdf"`               //固定值，传true
	IsLocalOrNetworkPdf int64  `json:"isLocalOrNetworkPdf"` //固定值 传0
	PdfUrl              string `json:"pdfUrl"`              // 地址
	QyId                string `json:"qyId"`                //企业编号
	File                []byte `json:"file"`                //文件二进制流
	EmpCode             string `json:"empCode"`             //用户ID
}

/**
获取微信Access_tokenTicket出参
*/
type Access_tokenTicket struct {
	TokenTicketUrl string `json:"tokenTicketUrl"`
	Signature      string `json:"signature"`
	AppId          string `json:"appId"`
	NonceStr       string `json:"nonceStr"`
	Timestamp      string `json:"timestamp"`
}
