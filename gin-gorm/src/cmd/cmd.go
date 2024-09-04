package cmd

import (
	"bronya.com/gin-gorm/src/conf"
	"bronya.com/gin-gorm/src/router"
	"log"
)

func Start() {
	conf.Init()
	router.InitRouter()
}

func Done() {
	log.Println("==================== Done! ====================")
}
