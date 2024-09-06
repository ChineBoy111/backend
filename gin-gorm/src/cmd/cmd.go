package cmd

import (
	"bronya.com/gin-gorm/src/conf"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/router"
	"context"
	"time"
)

func Start() {
	conf.ReadConf()                  //! 读取配置文件
	global.Logger = conf.NewLogger() //! 启动日志
	session, err := conf.ConnMysql() //! 连接 mysql，创建表

	if err != nil {
		global.Logger.Errorf("Connect mysql error %s", err.Error())
		panic(err.Error())
	}
	global.DB = session

	redisCli, err := conf.ConnRedis() //! 连接 redis，创建表
	if err != nil {
		global.Logger.Errorf("Connect redis error %s", err.Error())
		panic(err.Error())
	}
	global.RedisCli = redisCli

	err = global.RedisCli.Set(context.Background(), "username", "Tiancheng", 24*60*60*time.Second).Err()
	if err != nil {
		global.Logger.Errorf("Redis-cli set error %s", err.Error())
	}

	router.StartRouter() //! 注册路由，启动路由器
}

func Done() {
	global.Logger.Infoln("==================== Done! ====================")
}
