package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// IRoutesRegFunc 用于注册路由的函数
type IRoutesRegFunc = func(publicRouteGroup, authenRouteGroup *gin.RouterGroup)

var routersRegFuncArr []IRoutesRegFunc

// StartRouter 注册路由，启动路由器
func StartRouter() {

	//! 创建一个接收 os 信号的上下文 notifyCtx
	//! 收到任一 os 信号时，notifyCtx 的 Done 通道关闭，可执行 <-notifyCtx.Done()
	notifyCtx, notifyCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer notifyCancel()

	engine := gin.Default()
	publicRouteGroup := engine.Group("/api/v1/public")
	authenRouteGroup := engine.Group("/api/v1")

	//* 添加未注册的路由
	AppendUserRoutes()

	for _, routersRegFunc := range routersRegFuncArr {
		//* 注册路由
		routersRegFunc(publicRouteGroup, authenRouteGroup) // init
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	//! 在新协程中启动服务器，主协程不会阻塞，继续运行
	go func() {
		log.Printf("Serving on port %s\n", port)
		//! 不建议使用 err != http.ErrServerClosed
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Serve error %s\n", err.Error())
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
		log.Printf("Shutdown error %s\n", err.Error())
	}
	// <-timoutCtx.Done()
	log.Println("Shutdown ok")
}

// Appender 添加未注册的路由
func Appender(routersRegFunc IRoutesRegFunc) {
	if routersRegFunc == nil {
		return
	}
	routersRegFuncArr = append(routersRegFuncArr, routersRegFunc)
}
