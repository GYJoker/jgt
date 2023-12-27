package req

import (
	"encoding/json"
	"github.com/GYJoker/jgt/constants"
	"github.com/labstack/echo/v4"
)

// SetReqParamInfo 设置请求参数信息
func SetReqParamInfo(c echo.Context, param interface{}) {
	c.Set(constants.ReqParamKey, param)
}

// GetReqParamInfo 获取请求参数信息
func GetReqParamInfo(c echo.Context) interface{} {
	return c.Get(constants.ReqParamKey)
}

// SetTraceId 设置链路追踪ID
func SetTraceId(c echo.Context, traceId string) {
	if len(GetTraceId(c)) != 0 {
		return
	}
	c.Request().Header.Set(constants.TraceIdKey, traceId)
}

// GetTraceId 获取链路追踪ID
func GetTraceId(c echo.Context) string {
	return c.Request().Header.Get(constants.TraceIdKey)
}

// SetMuddleName 设置模块名称
func SetMuddleName(c echo.Context, muddleName string) {
	c.Request().Header.Set(constants.MuddleName, muddleName)
}

// GetMuddleName 获取模块名称
func GetMuddleName(c echo.Context) string {
	return c.Request().Header.Get(constants.MuddleName)
}

// GetJwtToken 获取JWT Token
func GetJwtToken(c echo.Context) string {
	return c.Request().Header.Get(constants.JwtTokenKey)
}

// SetUserInfo 设置用户信息
func SetUserInfo(c echo.Context, userInfo interface{}) {
	c.Set(constants.UserInfoKey, userInfo)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c echo.Context) uint64 {
	return c.Get(constants.UserInfoKey).(uint64)
}

// SetUserId 设置用户ID
func SetUserId(c echo.Context, userId uint64) {
	c.Set(constants.UserIdKey, userId)
}

// GetUserId 获取用户ID
func GetUserId(c echo.Context) uint64 {
	return getContextUint64Value(c, constants.UserIdKey)
}

func SetUserRoles(c echo.Context, userRole []*UserRoleDto) {
	permissionStr, _ := json.Marshal(userRole)
	c.Set(constants.UserRolesKey, string(permissionStr))
}

func GetUserRoles(c echo.Context) []*UserRoleDto {
	value := c.Get(constants.UserRolesKey)
	if value == nil {
		return make([]*UserRoleDto, 0)
	}

	permissionStr := value.(string)
	var permissionList []*UserRoleDto
	_ = json.Unmarshal([]byte(permissionStr), &permissionList)
	return permissionList
}

func SetUserName(c echo.Context, userName string) {
	c.Set(constants.UserNameKey, userName)
}

func GetUserName(c echo.Context) string {
	return getContextStringValue(c, constants.UserNameKey)
}

func SetRefreshToken(c echo.Context, refreshToken string) {
	c.Set(constants.RefreshTokenKey, refreshToken)
}

func SetMerchantCode(c echo.Context, merchantCode string) {
	c.Set(constants.MerchantCodeKey, merchantCode)
}

func SetCustomValue(c echo.Context, key string, value interface{}) {
	c.Set(key, value)
}

func GetCustomValue(c echo.Context, key string) interface{} {
	return c.Get(key)
}

func GetMerchantCode(c echo.Context) string {
	return getContextStringValue(c, constants.MerchantCodeKey)
}

func GetRefreshToken(c echo.Context) string {
	return getContextStringValue(c, constants.RefreshTokenKey)
}

func GetAppId(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderAppIdKey)
}

func GetPlatform(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderPlatformKey)
}

func getContextStringValue(c echo.Context, key string) string {
	if get := c.Get(key); get != nil {
		return get.(string)
	}
	return ""
}

func getContextUint64Value(c echo.Context, key string) uint64 {
	if get := c.Get(key); get != nil {
		return get.(uint64)
	}
	return 0
}

func getContextBoolValue(c echo.Context, key string) bool {
	if get := c.Get(key); get != nil {
		return get.(bool)
	}
	return false
}
