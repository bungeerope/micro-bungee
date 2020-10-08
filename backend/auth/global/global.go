package global

import (
	"github.com/bungeerope/micro-bungee/conf"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB     *gorm.DB
	GLOBAL_CACHE  *redis.Client
	GLOBAL_CONF   conf.Server
	GLOBAL_VIP    *viper.Viper
	GLOBAL_LOGGER *zap.Logger
)
