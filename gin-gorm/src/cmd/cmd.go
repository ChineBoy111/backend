package cmd

import (
	"bronya.com/gin-gorm/src/conf"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/router"
)

func Start() {
	conf.ReadConf()                  //! 读取配置文件
	global.Logger = conf.NewLogger() //! 启动日志
	dbSession, err := conf.ConnDB()  //! 连接数据库，创建表
	if err != nil {
		global.Logger.Errorf("Connect database error %s", err.Error())
		panic(err.Error())
	}
	global.DBSession = dbSession
	router.StartRouter() //! 注册路由，启动路由器
}

func Done() {
	global.Logger.Infoln("==================== Done! ====================")
}
