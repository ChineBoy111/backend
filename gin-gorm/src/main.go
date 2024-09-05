package main

import (
	"bronya.com/gin-gorm/src/cmd"
	"os"
	"os/signal"
	"syscall"
)

func unused() {
	defer cmd.Done()

	//! 在新协程中启动服务器，主协程不会阻塞，继续运行
	go func() {
		cmd.Start()
	}()

	quitChan := make(chan os.Signal, 1)
	//? kill    发送 syscall.SIGTERM
	//? kill -2 发送 syscall.SIGINT (os.Interrupt)
	//? kill -9 发送 syscall.SIGKILL 但不能被捕获

	//! 将指定的 os 信号转发到 quitChan 通道
	signal.Notify(quitChan, syscall.SIGINT /* os.Interrupt */, syscall.SIGTERM)
	<-quitChan
}

func main() {
	defer cmd.Done()
	cmd.Start()
}
