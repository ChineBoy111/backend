package router

import (
	"bronya.com/gin-gorm/src/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRouteGroup 创建 user 子路由组
func UserRouteGroup() {
	userApi := api.NewUserApi()

	//* 创建 public 子路由组 publicUserRouteGroup
	publicUserRouteGroup := publicRouteGroup.Group("/user" /* 路由前缀 */)
	{ // publicUserRouteGroup
		publicUserRouteGroup.POST("/login", userApi.Login /* 回调函数 */)
	}

	//* 创建 authorized 子路由组 authorizedUserRouteGroup
	authorizedUserRouteGroup := authorizedRouteGroup.Group("/user" /* 路由前缀 */)
	{ // authorizedUserRouteGroup
		authorizedUserRouteGroup.GET("" /* 路由前缀 */, func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "One.c"},
					{"id": 1, "name": "Two.go"},
					{"id": 3, "name": "Three.ts"},
				},
			})
		})

		authorizedUserRouteGroup.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   4,
				"name": "Four.py",
			})
		})
	}
}
