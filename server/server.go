package server

import (
	"encoding/json"
	"fmt"
	"github.com/GYJoker/jgt/base"
	"github.com/GYJoker/jgt/cache"
	"github.com/GYJoker/jgt/config"
	"github.com/GYJoker/jgt/event_bus"
	"github.com/GYJoker/jgt/glog"
	"github.com/GYJoker/jgt/msg_nsq"
	"github.com/GYJoker/jgt/req"
	"github.com/GYJoker/jgt/resp"
	"github.com/GYJoker/jgt/yj_constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"net/http"
	"strings"
	"time"
)

type (
	Server struct {
		Version *base.Version

		isInitSuccess bool

		cc *config.Config

		// echo
		ec *echo.Echo

		// 数据库
		db *gorm.DB

		// 缓存redis
		redis cache.RedisManager

		// bus 全局事件总线
		bus event_bus.Bus

		// 消息中间件
		msgNsq msg_nsq.Manager
	}
)

func InitServerByConfig(conf *config.Config, version *base.Version) *Server {
	server := &Server{
		Version: version,
		cc:      conf,
		bus:     event_bus.New(), // 全局事件总线.
	}

	// 链接数据库
	sqlConnStr := server.cc.GetConnStr()
	if sqlConnStr != "" {
		d, e := gorm.Open(mysql.Open(sqlConnStr), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if e != nil {
			panic("db connect err" + e.Error())
			return nil
		}
		server.db = d
	}

	// 链接redis
	if server.cc.Redis != nil {
		server.redis = cache.NewManager(&cache.RedisConnOpt{
			Host:     server.cc.Redis.Host,
			Port:     server.cc.Redis.Port,
			Password: server.cc.Redis.Password,
		})
	}

	// 链接nsq
	if server.cc.Nsq != nil {
		server.msgNsq = msg_nsq.NewManager(server.cc.Nsq)
	}

	// 创建echo
	server.ec = echo.New()

	server.isInitSuccess = true

	// 添加路由
	server.addRouter()

	// 添加中间件
	server.addMiddleware()

	// 添加日志
	server.addLogger()

	return server
}

func InitServer(configId, configPath string, version *base.Version) *Server {
	if configPath != "" {
		config.UpdateConfigPath(configPath)
	}

	// 获取配置信息
	s, err := config.GetConfig(configId)
	if err != nil {
		panic("get config err: " + err.Error())
		return nil
	}

	return InitServerByConfig(s, version)
}

func (s *Server) StartServer() {
	// 启动定时删除临时文件
	//utils.TimerDeleteTempFile()

	// 启动服务
	s.ec.Logger.Fatalf(s.ec.Start(s.cc.ServerAddr()).Error())
}

func (s *Server) addRouter() {
	// 添加路由
	s.ec.GET("/", func(c echo.Context) error {
		return c.String(200, "hello world -- "+s.cc.Server.Label)
	})

	s.ec.GET("/ping", func(c echo.Context) error {
		return resp.ResponseBody(c, resp.GenSuccess("pong"))
	})
}

func (s *Server) addMiddleware() {
	// 添加中间件

	// 错误处理
	s.ec.HTTPErrorHandler = func(err error, c echo.Context) {
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
	s.ec.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req.SetMuddleName(c, s.cc.Server.Name)
			// 在这里进行请求预处理
			return next(c)
		}
	})

	// 允许跨域
	s.ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // "http://localhost:7200", "https://gongyj.net"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderContentType, "authorization", yj_constants.HeaderAppIdKey, yj_constants.HeaderVersionKey,
			yj_constants.HeaderMacKey, yj_constants.HeaderSignKey, yj_constants.HeaderPlatformKey, yj_constants.HeaderDeviceKey,
			yj_constants.HeaderTimestampKey, "fileName"},
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

func (s *Server) addLogger() {
	// 添加日志
}

func (s *Server) GetEcho() *echo.Echo {
	if !s.isInitSuccess || s.ec == nil {
		panic("server not init")
	}
	return s.ec
}

func (s *Server) GetDB() *gorm.DB {
	if !s.isInitSuccess || s.db == nil {
		panic("db not init")
	}
	return s.db
}

func (s *Server) GetRedis() cache.RedisManager {
	if !s.isInitSuccess || s.redis == nil {
		panic("redis not init")
	}
	return s.redis
}

func (s *Server) GetBus() event_bus.Bus {
	if !s.isInitSuccess || s.bus == nil {
		panic("bus not init")
	}
	return s.bus
}

func (s *Server) GetMsgNsq() msg_nsq.Manager {
	if !s.isInitSuccess || s.msgNsq == nil {
		panic("msg_nsq not init")
	}
	return s.msgNsq
}

func (s *Server) GetConfig() *config.Config {
	if !s.isInitSuccess || s.cc == nil {
		panic("server not init")
	}
	return s.cc
}

// ReportServerInfo 上报服务信息
func (s *Server) ReportServerInfo() {
	if strings.Contains(s.cc.ConfigId, "center_server") {
		return
	}
	server := s.cc.Server
	report, err := config.GetConfig("center_server")
	register := report.Server
	param := make(map[string]interface{})
	param["name"] = server.Name
	param["ip"] = server.LocalIp
	param["port"] = server.Port
	body, err := postReq("http://"+register.Host+":"+register.Port+"/"+yj_constants.RegisterPath, param)
	if err != nil {
		fmt.Println("report service info failed, g_error:", err)

		time.AfterFunc(time.Minute*2, s.ReportServerInfo)

		return
	}
	fmt.Println(body)
}

// PostReq post请求
func postReq(url string, param interface{}) (*resp.Body, error) {
	bytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(bytes))
	// 发送请求
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		return nil, err
	}
	return handleRespData(response)
}

// handleRespData 处理返回数据
func handleRespData(response *http.Response) (*resp.Body, error) {
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			glog.GetLogger().Error("close body failed, err:", err)
		}
	}(response.Body)
	// 解析返回数据
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		glog.GetLogger().Error("read body failed, err:", err)
		return nil, err
	}
	data := &resp.Body{}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return &resp.Body{
			Data: string(bytes),
			Code: response.StatusCode,
			Msg:  response.Status,
		}, nil
	}
	return data, nil
}
