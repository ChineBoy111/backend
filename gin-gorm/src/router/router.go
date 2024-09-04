package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// IRoutesRegFunc
// ! Function for registering routes
type IRoutesRegFunc = func(publicRouteGroup, authenRouteGroup *gin.RouterGroup)

var routersRegFuncArr []IRoutesRegFunc

// InitRouter
// ! Init router and start server
func InitRouter() {
	engine := gin.Default()
	publicRouteGroup := engine.Group("/api/v1/public")
	authenRouteGroup := engine.Group("/api/v1")

	appendRoutes()

	for _, routersRegFunc := range routersRegFuncArr {
		// ! Register routes
		routersRegFunc(publicRouteGroup, authenRouteGroup) // init
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	err := engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf("Start starded error %s", err.Error()))
	}
}

// AppendRoutes
// ! Append unregistered routes
func appendRoutes() {
	AppendUserRoutes()
}

// RoutesAppender
// ! Unregistered routes appender
func RoutesAppender(routersRegFunc IRoutesRegFunc) {
	if routersRegFunc == nil {
		return
	}
	routersRegFuncArr = append(routersRegFuncArr, routersRegFunc)
}
