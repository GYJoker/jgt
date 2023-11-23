package base

// Dto 请求通用参数
type Dto struct {
	AppId    string `json:"app_id" query:"app_id" form:"app_id" header:"app_id"`
	Version  string `json:"version,omitempty" form:"version" query:"version" header:"version"`
	Platform string `json:"platform,omitempty" form:"platform" query:"platform" header:"platform"`
	Device   string `json:"device,omitempty" form:"device" query:"device" header:"device"`
	Mac      string `json:"mac,omitempty" form:"mac" query:"mac" header:"mac"`
}
