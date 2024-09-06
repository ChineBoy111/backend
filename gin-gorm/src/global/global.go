package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//  全局变量

var (
	Logger   *zap.SugaredLogger
	DB       *gorm.DB
	RedisCli *redis.Client
)
