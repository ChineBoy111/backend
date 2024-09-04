package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"bronya.com/gin-gorm/src/cmd"
)

func main_() {
	defer cmd.Done()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown below 主协程可以继续运行
	go func() {
		cmd.Start()
	}()

	quitChan := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT, also os.Interrupt
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it

	//! 将指定的 os 信号转发到 quitChan 通道
	signal.Notify(quitChan, syscall.SIGINT /* os.Interrupt */, syscall.SIGTERM)
	<-quitChan
}

func main() {
	defer cmd.Done()
	rootCtx := context.Background()

	listenerCtx, cancelFunc := signal.NotifyContext(
		rootCtx, syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown below 主协程可以继续运行
	go func() {
		cmd.Start()
	}()

	<-listenerCtx.Done() // <-listenerCtx.Done() == {}
}
