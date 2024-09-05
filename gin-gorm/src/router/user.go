package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppendUserRoutes 添加未注册的 user 路由
func AppendUserRoutes() {
	Appender(func(publicRouteGroup, authenRouteGroup *gin.RouterGroup) {

		//* publicRouteGroup
		publicRouteGroup.POST("/login", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK /* 200 */, gin.H{
				"msg": "Login OK",
			} /* gin.H 是 map[string]any 的别名 */)
		} /* 回调函数 */)

		//* 创建一个 authenRouteGroup 的子路由组 userRouteGroup
		userRouteGroup := authenRouteGroup.Group("/user" /* 路由前缀 */)

		userRouteGroup.GET("" /* 路由前缀 */, func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "One.c"},
					{"id": 1, "name": "Two.go"},
					{"id": 3, "name": "Three.ts"},
				},
			})
		})

		userRouteGroup.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   4,
				"name": "Four.py",
			})
		})
	})
}
