package server

import (
	"fmt"
	"github.com/GYJoker/jgt/cache"
	"github.com/GYJoker/jgt/config"
	"github.com/GYJoker/jgt/constants"
	"github.com/GYJoker/jgt/event_bus"
	"github.com/GYJoker/jgt/req"
	"github.com/GYJoker/jgt/resp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
)

var (
	isInitSuccess = false

	// 服务配置
	cc *config.Config

	// echo
	ec *echo.Echo

	// 数据库
	db *gorm.DB

	// 缓存redis
	redis cache.RedisManager

	// bus 全局事件总线
	bus = event_bus.New()
)

func InitServer(configId, configPath string) {
	if configPath != "" {
		config.UpdateConfigPath(configPath)
	}

	// 获取配置信息
	s, err := config.GetConfig(configId)
	if err != nil {
		panic("get config err: " + err.Error())
		return
	}

	cc = s

	// 链接数据库
	d, err := gorm.Open(mysql.Open(cc.GetConnStr()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("db connect err" + err.Error())
		return
	}
	db = d

	// 链接redis
	if cc.Redis != nil {
		redis = cache.NewManager(&cache.RedisConnOpt{
			Host:     cc.Redis.Host,
			Port:     cc.Redis.Port,
			Password: cc.Redis.Password,
		})
	}

	// 创建echo
	ec = echo.New()

	isInitSuccess = true

	// 添加路由
	addRouter()

	// 添加中间件
	addMiddleware()

	// 添加日志
	addLogger()
}

func StartServer() {
	// 启动定时删除临时文件
	//utils.TimerDeleteTempFile()

	// 启动服务
	ec.Logger.Fatalf(ec.Start(cc.ServerAddr()).Error())
}

func addRouter() {
	// 添加路由
	ec.GET("/", func(c echo.Context) error {
		return c.String(200, "hello world -- "+cc.Server.Label)
	})

	ec.GET("/ping", func(c echo.Context) error {
		return resp.ResponseBody(c, resp.GenSuccess("pong"))
	})
}

func addMiddleware() {
	// 添加中间件

	// 错误处理
	ec.HTTPErrorHandler = func(err error, c echo.Context) {
		if err == nil {
			return
		}
		fmt.Println("http error: ", err.Error())
		if he, ok := err.(*echo.HTTPError); ok {
			_ = c.JSON(http.StatusOK, resp.GenError(he.Code, he.Error()))
			return
		}
		_ = c.JSON(http.StatusOK, resp.GenError(http.StatusInternalServerError, err.Error()))
	}

	// 请求预处理
	ec.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req.SetMuddleName(c, cc.Server.AppID)
			// 在这里进行请求预处理
			return next(c)
		}
	})

	// 允许跨域
	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // "http://localhost:7200", "https://gongyj.net"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderContentType, "authorization", constants.HeaderAppIdKey, constants.HeaderVersionKey,
			constants.HeaderMacKey, constants.HeaderSignKey, constants.HeaderPlatformKey, constants.HeaderDeviceKey,
			constants.HeaderTimestampKey, "fileName"},
		ExposeHeaders: []string{echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowOrigin, "fileName"},
	}))

	//defaultBodyDumpConfig := middleware.BodyDumpConfig{
	//	Skipper: bodyDumpDefaultSkipper,
	//	Handler: func(c echo.Context, reqBody []byte, resBody []byte) {
	//		println("API请求结果拦截：", string(resBody))
	//		// 1、解析返回的json数据，判断接口执行成功或失败。如： {"code":"200","data":"test","msg":"请求成功"}
	//		// 2、保存操作日志
	//	},
	//}
	//
	//ec.Use(middleware.BodyDumpWithConfig(defaultBodyDumpConfig))
}

func addLogger() {
	// 添加日志
}

func GetEcho() *echo.Echo {
	if !isInitSuccess || ec == nil {
		panic("server not init")
	}
	return ec
}

func GetDB() *gorm.DB {
	if !isInitSuccess || db == nil {
		panic("server not init")
	}
	return db
}

func GetRedis() cache.RedisManager {
	if !isInitSuccess || redis == nil {
		panic("server not init")
	}
	return redis
}

func GetBus() event_bus.Bus {
	if !isInitSuccess || bus == nil {
		panic("server not init")
	}
	return bus
}
