package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
	Server ServerConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

type ServerConfig struct {
	Ip string
	Port string
}
var Cfg *tomlConfig

//程序启动时执行init方法
func init() {
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "Go-blog"
	Cfg.System.Version = 1.0
	curDir, _ := os.Getwd()
	Cfg.System.CurrentDir = curDir
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}

}
