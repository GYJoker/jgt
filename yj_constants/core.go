package yj_constants

const (
	// SmsBitCount 短信位数
	SmsBitCount = 6
	// SmsValidTime 短信有效时间 单位分钟
	SmsValidTime = 5
	// SmsOneDayMaxCount 同一天短信最大发送次数
	SmsOneDayMaxCount = 10
	// SmsOneTypeMinMinute 同一类型短信最小间隔分钟数
	SmsOneTypeMinMinute = 1
	// SmsMaxFailedCount 验证失败次数
	SmsMaxFailedCount = 3
)

// 短信验证码类型
const (
	// SmsTypeLogin 登录短信验证码
	SmsTypeLogin = "login"
)

const (
	VipTypeDay   = "day"
	VipTypeMonth = "month"
	VipTypeYear  = "year"
)

// 系统配置类型
const (
	// ConfigTypeText 文本
	ConfigTypeText = "text"
	// ConfigTypeImage 图片
	ConfigTypeImage = "image"
	// ConfigTypeUrl 链接
	ConfigTypeUrl = "url"
	// ConfigTypeSwitch 开关
	ConfigTypeSwitch = "switch"
	// ConfigTypeNumber 数字
	ConfigTypeNumber = "number"
)

const (
	FeedbackTypeBug = "bug"
	FeedbackTypeSug = "sug" // 建议 suggest
)

const (
	FeedbackStatusNew    = "new"
	FeedbackStatusSure   = "sure"
	FeedbackStatusReject = "reject"
)

const (
	ThirdAppTypeWechatMiniApp = "wechat_mini_app"
	ThirdAppTypeWechatH5      = "wechat_h5"
	ThirdAppTypeAlipayMiniApp = "alipay_mini_app"
	ThirdAppTypeAlipayH5      = "alipay_h5"
	ThirdAppTypeAppWechat     = "app_wechat"
	ThirdAppTypeAppAlipay     = "app_alipay"
)

var (
	// ThirdAppTypeList 第三方应用类型列表
	ThirdAppTypeList = []string{
		ThirdAppTypeWechatMiniApp,
		ThirdAppTypeWechatH5,
		ThirdAppTypeAlipayMiniApp,
		ThirdAppTypeAlipayH5,
		ThirdAppTypeAppWechat,
		ThirdAppTypeAppAlipay,
	}
)
