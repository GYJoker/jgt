package resp

type ErrCode struct {
	Code int
	Msg  string
}

var (
	// SUCCESS 成功
	SUCCESS = &ErrCode{
		Code: SuccessCode,
		Msg:  "success",
	}

	// ParamsErr 参数错误
	ParamsErr = &ErrCode{
		Code: ParamErrCode,
		Msg:  "参数错误",
	}

	// CustomErr 自定义错误
	CustomErr = &ErrCode{
		Code: CustomErrCode,
		Msg:  "error",
	}

	// NotAuthErr 未授权
	NotAuthErr = &ErrCode{
		Code: NotAuthCode,
		Msg:  "未授权",
	}

	// RecordNotExistErr 找不到
	RecordNotExistErr = &ErrCode{
		Code: RecordNotExistCode,
		Msg:  "记录不存在",
	}

	// PermissionDeniedErr 权限不足
	PermissionDeniedErr = &ErrCode{
		Code: PermissionDeniedCode,
		Msg:  "权限不足",
	}

	AppErr = &ErrCode{
		Code: AppErrCode,
		Msg:  "APP权限受限",
	}

	SignErr = &ErrCode{
		Code: SignErrCode,
		Msg:  "签名错误",
	}
)

// 定义普通错误code
const (
	SuccessCode        = 100000
	ParamErrCode       = 100001
	CustomErrCode      = 100002
	SqlErrCode         = 100003
	RecordNotExistCode = 100004
	RecordExistCode    = 100005
	FileSystemErrCode  = 100006

	NotAuthCode          = 100099
	PermissionDeniedCode = 100100
	NormalErrCode        = 100101
	IPErrCode            = 100102
	AppErrCode           = 100103
	SignErrCode          = 100104
)

// 用户系统相关错误code
const (
	PwdErrCode     = 110003
	PhoneExistCode = 110004
)

func (e *ErrCode) SetCode(code int) *ErrCode {
	e.Code = code
	return e
}

func (e *ErrCode) SetMsg(msg string) *ErrCode {
	e.Msg = msg
	return e
}
