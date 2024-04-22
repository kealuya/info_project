package entity

//base struct
type Base struct {
	CorpName   string `json:"corpName" description:"企业名称" `
	CorpCode   string `json:"corpCode" description:"企业code"`
	CreateTime string `json:"createTime" description:"创建时间"`
	Creater    string `json:"creater" description:"创建人"`
	Bz1        string `json:"bz1" description:"备注1"`
	Bz2        string `json:"bz2" description:"备注2"`
	Bz3        string `json:"bz3" description:"备注3"`
}

//分页信息
type Paging struct {
	CurrentPage int64 `json:"currentPage" description:"当前页" `
	PageSize    int64 `json:"pageSize" description:"一页几条" `
	TotalCount  int64 `json:"totalCount" description:"总条数" `
	PageCount   int64 `json:"pageCount" description:"总页数" `
}
