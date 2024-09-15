package utils

import (
	"encoding/json"
	"log"
	"os"
)

type IGlobal struct {
	Name    string // 服务器名
	Ver     string // 服务器版本
	BufSize int    // 缓冲区大小
	HostIp  string // 监听的 ip 地址
	MaxConn int    // 最大连接数
	Port    int    // 监听的端口
	Proto   string // 协议
}

var Global *IGlobal

// ! init 函数只执行 1 次
func init() {
	Global = &IGlobal{
		Name:    "Proxy",
		Ver:     "1.0",
		BufSize: 512,
		HostIp:  "0.0.0.0",
		MaxConn: 1,
		Port:    8080,
		Proto:   "tcp4",
	}
	Global.Load()
}

func (*IGlobal) Load() {
	byteArr, err := os.ReadFile("./proxy.json")
	if err != nil {
		log.Println("Read file err", err.Error())
		return
	}
	// 解析 json 数据到 go 结构体变量
	err = json.Unmarshal(byteArr, Global)
}
