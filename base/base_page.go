package base

type (
	QueryPage struct {
		Dto
		Index int `json:"index" form:"index" query:"index"`
		Size  int `json:"size" form:"size" query:"size"`
	}

	RespPage struct {
		Current int         `json:"current"`  // 当前页
		Pages   int         `json:"pages"`    // 总页数
		Size    int         `json:"size"`     // 每页数量
		Total   int         `json:"total"`    // 总数量
		HasNext bool        `json:"has_next"` // 是否有下一页
		Records interface{} `json:"records"`  // 数据
	}
)

func (page *QueryPage) GetIndex() int {
	if page.Index <= 0 {
		return 1
	}
	return page.Index
}

func (page *QueryPage) GetSize() int {
	if page.Size <= 0 {
		return 10
	}
	return page.Size
}

func (page *QueryPage) GetOffset() int {
	return (page.GetIndex() - 1) * page.GetSize()
}

func (page *QueryPage) GetLimit() int {
	return page.GetSize()
}

func GetNilRespPage() *RespPage {
	return &RespPage{
		Current: 0,
		Pages:   0,
		Size:    0,
		Total:   0,
		Records: []interface{}{},
	}
}

func (resp *RespPage) UpdateData(data interface{}) {
	resp.Records = data
}

func (resp *RespPage) UpdatePage(total int64, size, index int) {
	pages := total / int64(size)
	if total%int64(size) > 0 {
		pages++
	}

	resp.Total = int(total)
	resp.Size = size
	resp.Pages = int(pages)
	resp.Current = index
	resp.HasNext = index < int(pages)
}
