package router

import (
	"bronya.com/gin-gorm/src/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ! rg     | router group
// ! pubRg  | public router group
// ! authRg | authorized route group

// RegUserRg 注册 user 路由组
func RegUserRg() {
	userApi := controller.NewUserApi()

	//* 创建 pubRg 的子路由组 pubUserRg
	pubUserRg := pubRg.Group("/user" /* 路由前缀 */)
	{
		pubUserRg.POST("/login", userApi.Login /* 回调函数 */)
	}

	//* 创建 authRg 的子路由组 authUserRg
	authUserRg := authRg.Group("/user" /* 路由前缀 */)
	{ // authUserRg
		authUserRg.GET("" /* 路由前缀 */, func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "One.c"},
					{"id": 1, "name": "Two.go"},
					{"id": 3, "name": "Three.ts"},
				},
			})
		})

		authUserRg.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   4,
				"name": "Four.py",
			})
		})
	}
}
