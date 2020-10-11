package database

import (
	"github.com/bungeerope/micro-bungee/global"
	"github.com/bungeerope/micro-bungee/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var err error

func Gorm() {
	switch global.GLOBAL_CONF.System.DbType {
	case "mysql":
		initMysqlGorm()
	}
}

func initMysqlGorm() {
	m := global.GLOBAL_CONF.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Url + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	// 配置日志模式
	gormConfig := config(m.LogMode)
	if global.GLOBAL_DB, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		global.GLOBAL_LOGGER.Error("Mysql 连接异常", zap.Any("error", err))
		os.Exit(0)
	}
	GormTableRegister(global.GLOBAL_DB)
	mysqlDB, _ := global.GLOBAL_DB.DB()
	mysqlDB.SetMaxIdleConns(m.MaxIdleConns)
	mysqlDB.SetMaxOpenConns(m.MaxOpenConns)
}

// 注册数据库中相关表
func GormTableRegister(db *gorm.DB) {
	var err = db.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
	)
	if err != nil {
		global.GLOBAL_LOGGER.Error("register table failed", zap.Any("error", err))
		os.Exit(0)
	}
	global.GLOBAL_LOGGER.Info("register table success")

}

func config(logMode bool) (c *gorm.Config) {
	if logMode {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}
