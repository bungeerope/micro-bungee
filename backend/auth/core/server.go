package core

import "github.com/bungeerope/micro-bungee/global"

type server interface {
	Servers() error
}

func RunServer() {
	if global.GLOBAL_CONF.System.MultiPoint {
		// 初始化redis服务
		//initialize.Redis()
	}

}
