package router

import (
	"net/http"

	"bronya.com/gin-gorm/src/api"
	"github.com/gin-gonic/gin"
)

// AppendUserRoutes 添加未注册的 user 路由
func AppendUserRoutes() {
	userApi := api.NewUserApi()

	Appender(func(pubRouteGroup, authRouteGroup *gin.RouterGroup) {

		//* 创建一个 pubRouteGroup 的子路由组 pubUserRouteGroup
		pubUserRouteGroup := pubRouteGroup.Group("/user" /* 路由前缀 */)

		{ // pubUserRouteGroup
			pubUserRouteGroup.POST("/login", userApi.Login /* 回调函数 */)
		}

		//* 创建一个 authRouteGroup 的子路由组 authUserRouteGroup
		authUserRouteGroup := authRouteGroup.Group("/user" /* 路由前缀 */)

		{ // authUserRouteGroup
			authUserRouteGroup.GET("" /* 路由前缀 */, func(context *gin.Context) {
				context.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "One.c"},
						{"id": 1, "name": "Two.go"},
						{"id": 3, "name": "Three.ts"},
					},
				})
			})

			authUserRouteGroup.GET("/:id", func(context *gin.Context) {
				context.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   4,
					"name": "Four.py",
				})
			})
		}
	})
}
