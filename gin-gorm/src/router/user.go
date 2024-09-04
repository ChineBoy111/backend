package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppendUserRoutes
// ! Append unregistered routes for user
func AppendUserRoutes() {
	RoutesAppender(func(publicRouteGroup, authenRouteGroup *gin.RouterGroup) {

		//* publicRouteGroup
		publicRouteGroup.POST("/login", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK /* 200 */, gin.H{
				"msg": "Login OK",
			} /* gin.H is a shortcut for map[string]any */)
		} /* callback */)

		//* creates a new sub route group of authenRouteGroup
		userRouteGroup := authenRouteGroup.Group("/user" /* relativePath (prefix) */)

		userRouteGroup.GET("" /* relativePath (prefix) */, func(context *gin.Context) {
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
