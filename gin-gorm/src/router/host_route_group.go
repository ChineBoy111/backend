package router

import "bronya.com/gin-gorm/src/api"

func HostRouteGroup() {
	hostApi := api.NewHostApi()

	//* 创建 authorizedRouteGroup 的子路由组 authorizedHostRouteGroup
	authorizedHostRouteGroup := authorizedRouteGroup.Group("host" /* 路由前缀 */)
	authorizedHostRouteGroup.GET("/shutdown", hostApi.Shutdown)
}
