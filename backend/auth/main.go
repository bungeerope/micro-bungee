package main

import (
	"github.com/bungeerope/micro-bungee/core"
	"github.com/bungeerope/micro-bungee/global"
	"github.com/bungeerope/micro-bungee/init/database"
)

func main() {
	// 初始化数据库
	database.Gorm()
	// 程序退出前关闭数据库连接
	db, _ := global.GLOBAL_DB.DB()
	defer db.Close()

	core.RunServer()
}
