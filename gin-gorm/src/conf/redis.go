package conf

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ConnRedis() (dbSession *gorm.DB, err error) {
	redisClient := redis.NewClient(&redis.Options{})
}
