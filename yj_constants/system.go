package yj_constants

const NormalUserRole = "normal"

const (
	DefaultPwd           = "123456"
	FileBaseUrl          = "https://file.jiujunet.com/"
	DefaultUserHeaderImg = "https://file.jiujunet.com/avatar/default_avatar.png"
)

const (
	ManagerTokenValidTime int64 = 60 * 60
	MiniAppTokenValidTime int64 = 60 * 60 * 24 * 15
)

const (
	OneDayInSeconds = 60 * 60 * 24
)

const (
	SystemTrueValue  = 1
	SystemFalseValue = 0
)

const (
	SystemRecordNoData = 0
	SystemRecordExist  = 1
)

// 平台类型
const (
	// PlatformTypeAndroid 安卓平台
	PlatformTypeAndroid = "android"
	// PlatformTypeIos IOS平台
	PlatformTypeIos = "ios"
	// PlatformTypeWeb WEB平台
	PlatformTypeWeb = "web"
	// PlatformTypeMiniAppWechatIos 微信小程序IOS平台
	PlatformTypeMiniAppWechatIos = "mini_app_wechat_ios"
	// PlatformTypeMiniAppWechatAndroid 微信小程序安卓平台
	PlatformTypeMiniAppWechatAndroid = "mini_app_wechat_android"
	// PlatformTypeMiniAppWechatDev 微信小程序开发平台
	PlatformTypeMiniAppWechatDev = "mini_app_wechat_dev"
	// PlatformTypeMiniAppAlipayIos 支付宝小程序IOS平台
	PlatformTypeMiniAppAlipayIos = "mini_app_alipay_ios"
	// PlatformTypeMiniAppAlipayAndroid 支付宝小程序安卓平台
	PlatformTypeMiniAppAlipayAndroid = "mini_app_alipay_android"
	// PlatformTypeMiniAppAlipayDev 支付宝小程序开发平台
	PlatformTypeMiniAppAlipayDev = "mini_app_alipay_dev"
	// PlatformTypeH5AliPay 支付宝H5平台
	PlatformTypeH5AliPay = "h5_alipay"
	// PlatformTypeH5Wechat 微信H5平台
	PlatformTypeH5Wechat = "h5_wechat"
)

var (
	PlatformList = []string{
		PlatformTypeWeb,
		PlatformTypeAndroid,
		PlatformTypeIos,
		PlatformTypeMiniAppWechatIos,
		PlatformTypeMiniAppWechatAndroid,
		PlatformTypeMiniAppWechatDev,
		PlatformTypeMiniAppAlipayIos,
		PlatformTypeMiniAppAlipayAndroid,
		PlatformTypeMiniAppAlipayDev,
		PlatformTypeH5AliPay,
		PlatformTypeH5Wechat,
	}

	PlatformMiniAppList = []string{
		PlatformTypeMiniAppWechatIos,
		PlatformTypeMiniAppWechatAndroid,
		PlatformTypeMiniAppWechatDev,
		PlatformTypeMiniAppAlipayIos,
		PlatformTypeMiniAppAlipayAndroid,
		PlatformTypeMiniAppAlipayDev,
	}
)

const (
	IntMoneyRate    = 100
	DoubleMoneyRate = 100.0
)

const (
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionCreate = "create"
)

const (
	RegisterPath         = "register"
	CenterServerName     = "center"
	CustomerServerName   = "customer"
	UserServerName       = "user"
	CoreServerName       = "core"
	StoreServerName      = "store"
	LetterServerName     = "letter"
	RecordServerName     = "record"
	PermissionServerName = "permission"
	OrderServerName      = "order"
	PayServerName        = "pay"
	BeggingServerName    = "begging"
	QrCodeServerName     = "qr_code"
	RemindServerName     = "remind"
	GuessHeroServerName  = "guess_hero"
)

// 记录状态
const (
	RecordStatusUsed   = "used"
	RecordStatusUnused = "unused"
)

// NotFoundIndex 当查询不到数据时的默认Index
const NotFoundIndex = -1

const (
	// TagsMinCount 最小标签数量
	TagsMinCount = 3
	// TagsMaxCount 最大标签数量
	TagsMaxCount = 20
)
