package database

import "github.com/bungeerope/micro-bungee/global"

func Gorm() {
	switch global.GLOBAL_CONF.System.DbType {
	case "mysql":
		initMysqlGorm()
	}
}

func initMysqlGorm() {

}
