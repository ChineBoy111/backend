package network

import (
	"bronya.com/net-proxy/inetwork"
)

// TcpBaseMidWare tcp 基础中间件
type TcpBaseMidWare struct {
}

// PreHandler MsgHandler 前的 hook 方法
func (midWare *TcpBaseMidWare) PreHandler(inetwork.ITcpReq) {
}

// MsgHandler 处理拆包得到的 tcp 消息
func (midWare *TcpBaseMidWare) MsgHandler(inetwork.ITcpReq) {
}

// PostHandler MsgHandler 后的 hook 方法
func (midWare *TcpBaseMidWare) PostHandler(inetwork.ITcpReq) {
}
