package test

import (
	"bronya.com/gin-gorm/src/conf"
	"bronya.com/gin-gorm/src/util"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
	"time"
)

// ! 从一个已关闭的空通道中读，返回通道元素类型的零值和 false（表示读失败）
// ! go test -run TestNotify
func TestNotify(t *testing.T) {
	ch := make(chan int) // make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()
	fmt.Println(<-ch) // 0
}

// ! go test -run TestRedisCli
func TestRedisCli(t *testing.T) {
	redisCli, err := conf.ConnRedis() //! 连接 redis，创建表
	if err != nil {
		panic(fmt.Sprintf("Connect redis server error %s", err.Error()))
	}

	expiration := viper.GetDuration("db.redis.expiration") // 10 分钟

	//! redisCli.Set(context.Background(), key, value, expiration).Err()
	err = redisCli.Set(context.Background(), "username", "root", expiration*time.Second).Err()
	if err != nil {
		log.Printf("Redis set error %s\n", err.Error())
	}

	//! redisCli.Set(context.Background(), key, value, expiration).Err()
	err = redisCli.Set(context.Background(), "password", "0228", expiration*time.Second).Err()
	if err != nil {
		log.Printf("Redis set error %s\n", err.Error())
	}

	//! redisCli.Get(context.Background(), key).Result()
	username, err := redisCli.Get(context.Background(), "username").Result()
	if err != nil {
		log.Printf("Redis get error %s\n", err.Error())
	}
	log.Printf("Redis get username = %s", username)

	//! redisCli.Del(context.Background(), key1, key2, ...).Err()
	err = redisCli.Del(context.Background(), "username", "password").Err()
	if err != nil {
		log.Printf("Redis delete error %s\n", err.Error())
	}
}

// ! go test -run TestToken
func TestToken(t *testing.T) {
	viper.AddConfigPath("../")
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Read in config error %s", err.Error()))
	}

	log.Println("==================== Generate token ====================")
	tokStr, err := util.GenToken(1, "root")
	if err != nil {
		panic(fmt.Sprintf("Generate token error %s", err.Error()))
	}
	log.Printf("token = %s\n", tokStr)
	log.Println("==================== Parse token ====================")
	payload, err, isValid := util.ParseToken(tokStr) // tokStr + "suffix"
	if err != nil {
		panic(fmt.Sprintf("Parse token error %s, isValid = %v", err.Error(), isValid))
	}
	log.Printf("payload = %v, isValid = %v\n", payload, isValid)
}
