package flowablesdk

import "strings"

type Config struct {
	Url             string // 请求地址
	Username        string // 用户名
	Password        string // 密码
	ProcessPrefix   string // 流程API前缀
	FormPrefix      string // 表单API前缀
	ExtensionPrefix string // 自定义扩展API前缀
	RequestDebug    bool   // 是否打印请求日志
	ResponseDebug   bool   // 是否打印返回日志
}

var Configs Config

func Setup(c Config) {
	if len(c.Url) == 0 {
		c.Url = "http://127.0.0.1:8080"
	} else {
		c.Url = strings.Trim(c.Url, "/")
	}

	if len(c.ProcessPrefix) == 0 {
		c.ProcessPrefix = "/service"
	} else {
		c.ProcessPrefix = strings.Trim(c.ProcessPrefix, "/")
		c.ProcessPrefix = "/" + c.ProcessPrefix
	}

	if len(c.FormPrefix) == 0 {
		c.FormPrefix = "/form-api"
	} else {
		c.FormPrefix = strings.Trim(c.FormPrefix, "/")
		c.FormPrefix = "/" + c.FormPrefix
	}

	if len(c.ExtensionPrefix) == 0 {
		c.ExtensionPrefix = "/extension"
	} else {
		c.ExtensionPrefix = strings.Trim(c.ExtensionPrefix, "/")
		c.ExtensionPrefix = "/" + c.ExtensionPrefix
	}

	Configs.Url = c.Url
	Configs.Username = c.Username
	Configs.Password = c.Password
	Configs.ProcessPrefix = c.ProcessPrefix
	Configs.FormPrefix = c.FormPrefix
	Configs.ExtensionPrefix = c.ExtensionPrefix
	Configs.RequestDebug = c.RequestDebug
	Configs.ResponseDebug = c.ResponseDebug
}
