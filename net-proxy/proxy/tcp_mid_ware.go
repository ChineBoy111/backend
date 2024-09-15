package proxy

import (
	"bronya.com/net-proxy/iproxy"
)

type TcpBaseMidWare struct {
}

// PreHandler Handler 前的 hook 方法
func (midWare *TcpBaseMidWare) PreHandler(iproxy.ITcpReq) {
}

// MsgHandler 处理拆包得到的 tcp 消息
func (midWare *TcpBaseMidWare) MsgHandler(iproxy.ITcpReq) {
}

// PostHandler Handler 后的 hook 方法
func (midWare *TcpBaseMidWare) PostHandler(iproxy.ITcpReq) {
}
