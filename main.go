package main

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/server"
)

func init() {
	//模版加载
	common.LoadTemplate()
}

func main() {
	server.App.Start(config.Cfg.Server.Ip, config.Cfg.Server.Port)
}
