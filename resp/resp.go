package resp

import (
	"github.com/GYJoker/jgt/req"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"strings"
)

type Body struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id,omitempty"`
	Token     string      `json:"token,omitempty"`
}

type Log struct {
	RequestId  string      `json:"request_id"`
	Url        string      `json:"url"`
	MuddleName string      `json:"muddle_name"`
	Method     string      `json:"method"`
	UserId     uint64      `json:"user_id"`
	Param      interface{} `json:"param"`
	Resp       interface{} `json:"resp_body"`
	Token      string      `json:"token"`
}

type logFunc func(log *Log)

var logCallback logFunc

func SetLogFunc(f logFunc) {
	logCallback = f
}

func (b *Body) IsSuccess() bool {
	return b.Code == SUCCESS.Code
}

// GenSuccess 返回成功
func GenSuccess(data interface{}) *Body {
	return &Body{
		Code: SUCCESS.Code,
		Msg:  SUCCESS.Msg,
		Data: data,
	}
}

// GenError 返回错误
func GenError(code int, msg string) *Body {
	return &Body{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// GenErrorByCode 根据错误码返回错误
func GenErrorByCode(errCode *ErrCode) *Body {
	return &Body{
		Code: errCode.Code,
		Msg:  errCode.Msg,
		Data: nil,
	}
}

func ResponseFile(c echo.Context, body *Body) error {
	if !body.IsSuccess() {
		return ResponseBody(c, body)
	}
	fileName := body.Data.(string)
	split := strings.Split(fileName, "/")
	name := url.QueryEscape(split[len(split)-1])

	// 设置响应文件名
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+name)
	c.Response().Header().Set("fileName", name)

	return c.File(fileName)
}

func ResponseBody(c echo.Context, body *Body) error {
	path := c.Request().URL.String()
	if path == "/ping" {
		return c.JSON(http.StatusOK, body)
	}

	refreshToken := req.GetRefreshToken(c)
	if refreshToken != "" {
		body.Token = refreshToken
	}

	// 删除path中的mac参数
	if strings.Contains(path, "mac") {
		split := strings.Split(path, "mac=")
		path = split[0]
		last := split[1]
		if strings.Contains(last, "&") {
			split := strings.Split(last, "&")
			path += split[1]
		}
	}

	userId := req.GetUserId(c)
	body.RequestId = req.GetTraceId(c)
	log := &Log{
		RequestId:  body.RequestId,
		Token:      req.GetJwtToken(c),
		Url:        path,
		Param:      req.GetReqParamInfo(c),
		Resp:       body,
		Method:     c.Request().Method,
		UserId:     userId,
		MuddleName: req.GetMuddleName(c),
	}

	go writeLogInfo(log)

	return c.JSON(http.StatusOK, body)
}

func writeLogInfo(log *Log) {
	if logCallback != nil {
		go logCallback(log)
	}
}
