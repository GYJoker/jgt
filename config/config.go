package config

import (
	"fmt"
	"github.com/GYJoker/jgt/cron"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var (
	baseData           *BaseData
	configMap          = make(map[string]*Config)
	loadLock           sync.Mutex
	initOnce           sync.Once
	fileLastModifyTime int64
)

type (
	BaseData struct {
		Debug        bool           `yaml:"debug"`
		Encrypt      *EncryptConfig `yaml:"encrypt"`
		WebAppUid    string         `yaml:"web_app_uid"`
		Service      []*Config      `yaml:"service"`
		IpWhiteList  []string       `yaml:"ip_white_list"`
		PayNotifyUrl string         `yaml:"pay_notify_url"`
	}

	Config struct {
		// server配置
		Server *ServerConfig `json:"server" yaml:"server"`
		// 数据库配置
		MySql *DbConfig `json:"mysql" yaml:"mysql"`
		// redis配置
		Redis *RedisConfig `json:"redis" yaml:"redis"`
		// 服务名称
		ConfigId string `json:"config_id" yaml:"config_id"`
	}

	ServerConfig struct {
		Host           string         `json:"host"`
		Port           string         `json:"port"`
		LocalIp        string         `json:"local_ip" yaml:"local_ip"`
		Name           string         `json:"name"`
		Label          string         `json:"label"`
		StaticFilePath string         `json:"static_file_path" yaml:"static_file_path"`
		AppID          string         `json:"app_id" yaml:"app_id"`
		Encrypt        *EncryptConfig `json:"encrypt" yaml:"encrypt"`
	}

	DbConfig struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user_server"`
		Password string `json:"password"`
		Database string `json:"database" yaml:"database"`
	}

	RedisConfig struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
	}

	EncryptConfig struct {
		PwdSalt    string `json:"pwd_salt" yaml:"pwd_salt"`
		RsaPubFile string `json:"rsa_pub_file" yaml:"rsa_pub_file"`
		RsaPriFile string `json:"rsa_pri_file" yaml:"rsa_pri_file"`
	}
)

var configPath = "./config.yaml"

func UpdateConfigPath(path string) {
	configPath = path
}

// GetConfig 根据名称获取配置
func GetConfig(configId string) (*Config, error) {
	initOnce.Do(func() {
		loadConfigData()

		cronTime := &cron.YJCronTime{
			TimeType: cron.JobTimeSecond,
			Space:    30,
			Job: &cron.YJJob{
				JobFunc: func() {
					checkConfigFileInfo()
				},
			},
		}
		_, _ = cronTime.AddJob()
	})

	return configMap[configId], nil
}

func GetEncryptConfig() *EncryptConfig {
	return baseData.Encrypt
}

func GetPwdSalt() string {
	return baseData.Encrypt.PwdSalt
}

func GetPayNotifyUrl() string {
	return baseData.PayNotifyUrl
}

// GetMuddleMap 获取所有的muddle map
func GetMuddleMap() map[string]string {
	muddleMap := make(map[string]string)
	for _, v := range baseData.Service {
		muddleMap[v.Server.Name] = v.Server.Label
	}
	return muddleMap
}

// GetAdminWebAppId 获取后台管理系统的web app id
func GetAdminWebAppId() string {
	return baseData.WebAppUid
}

func (c *Config) ServerAddr() string {
	return ":" + c.Server.Port
}

// IsDebug 系统运行环境是否是debug模式
func IsDebug() bool {
	if baseData == nil {
		return true
	}
	return baseData.Debug
}

// IsInternalIp 是否是内部ip
func IsInternalIp(ip string) bool {
	if baseData == nil {
		return false
	}

	for _, v := range baseData.IpWhiteList {
		if v == ip {
			return true
		}
	}
	return false
}

func (c *Config) GetConnStr() string {
	if c.MySql == nil {
		return ""
	}
	return c.MySql.User +
		":" +
		c.MySql.Password +
		"@tcp(" + c.MySql.Host + ":" + c.MySql.Port + ")/" +
		c.MySql.Database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
}

// loadConfigData 加载配置文件
func loadConfigData() {
	loadLock.Lock()
	defer loadLock.Unlock()
	configMap = make(map[string]*Config)
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &baseData)
	if err != nil {
		return
	}

	for _, v := range baseData.Service {
		configMap[v.ConfigId] = v
	}
}

// checkConfigFileInfo 检查配置文件信息是否有修改
func checkConfigFileInfo() {
	file, err := os.OpenFile(configPath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("checkConfigFileInfo err:", err)
	}
	defer func(file *os.File) {
		e := file.Close()
		if e != nil {
			fmt.Println("checkConfigFileInfo close err:", err)
		}
	}(file)

	info, err := file.Stat()
	if err != nil {
		fmt.Println("checkConfigFileInfo stat err:", err)
	}

	if fileLastModifyTime == 0 {
		fileLastModifyTime = info.ModTime().Unix()
		return
	}

	if info.ModTime().Unix() == fileLastModifyTime {
		return
	}

	fmt.Println("checkConfigFileInfo modify reload")
	fileLastModifyTime = info.ModTime().Unix()
	loadConfigData()
}
