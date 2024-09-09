package router

import (
	"bronya.com/gin-gorm/src/api"
	_ "bronya.com/gin-gorm/src/docs"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

var publicRouteGroup, authorizedRouteGroup *gin.RouterGroup

// StartRouter 创建路由组（根路由组、子路由组），启动路由器
func StartRouter() {

	//! 创建一个接收 os 信号的上下文 notifyCtx
	//! 收到任一 os 信号时，notifyCtx 的 Done 通道关闭，可执行 <-notifyCtx.Done()
	notifyCtx, notifyCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer notifyCancel()

	//! 创建 gin 默认引擎
	engine := gin.Default()

	//! 创建根路由组
	publicRouteGroup = engine.Group("/api/v1/public")
	authorizedRouteGroup = engine.Group("/api/v1")
	//! 创建子路组
	UserRouteGroup() // 创建 user 子路由组

	//! 注册自定义字段校验器
	api.RegisterCustomValidator()

	//! 使用 gin 跨域中间件
	engine.Use(middleware.Cors())

	//* 访问 api 文档 http://127.0.0.1:3333/swagger/index.html
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := viper.GetString("server.port") // 3333
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	//! 在新协程中启动服务器，主协程不会阻塞，继续运行
	go func() {
		global.Logger.Infof("Serving on http://127.0.0.1:%s", port)
		//! 不建议使用 err != http.ErrServerClosed
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Errorf("Serve error %s", err.Error())
			return
		}
	}()

	<-notifyCtx.Done()

	//! 创建一个有超时时间的上下文 timeoutCtx
	//! 超时时间到时，timeoutCtx 的 Done 通道关闭，可执行 <-timoutCtx.Done()
	timoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeoutCancel()

	if err := server.Shutdown(timoutCtx); /* server.Shutdown(timeoutCtx) 会执行 <-timoutCtx.Done() */
	err != nil {
		global.Logger.Errorf("Shutdown error %s", err.Error())
	}
	// <-timoutCtx.Done()
	global.Logger.Infoln("Shutdown ok")
}
