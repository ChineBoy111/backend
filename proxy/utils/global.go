package utils

import (
	"bronya.com/proxy/iproxy_net"
	"encoding/json"
	"log"
	"os"
)

type IGlobal struct {
	Name          string // 服务器名
	Ver           string // 服务器版本
	Ip            string // 服务器监听的 ip 地址
	MaxConn       int    // 最大连接数
	MaxPacketSize int    // 最大数据包大小
	TcpPort       int    // tcp 服务器监听的端口
	TcpServer     iproxy_net.ITcpServer
}

var Global *IGlobal

// ! init 函数只会执行 1 次
func init() {
	Global = &IGlobal{
		Name:          "WAN Proxy",
		Ver:           "1.0",
		Ip:            "127.0.0.1",
		MaxConn:       100,
		MaxPacketSize: 512,
		TcpPort:       3333,
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
