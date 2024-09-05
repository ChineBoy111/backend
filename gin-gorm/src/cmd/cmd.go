package cmd

import (
	"bronya.com/gin-gorm/src/conf"
	"bronya.com/gin-gorm/src/router"
	"log"
)

func Start() {
	conf.LoadConf()      //! 加载配置文件
	router.StartRouter() //! 注册路由，启动路由器
}

func Done() {
	log.Println("==================== Done! ====================")
}
