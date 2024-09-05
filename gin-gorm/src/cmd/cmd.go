package cmd

import (
	"bronya.com/gin-gorm/src/conf"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/router"
	"log"
)

func Start() {
	conf.ReadConf()                  //! 读取配置文件
	global.Logger = conf.NewLogger() //! 启动日志
	router.StartRouter()             //! 注册路由，启动路由器
}

func Done() {
	log.Println("==================== Done! ====================")
	global.Logger.Infoln("==================== Done! ====================")
}
