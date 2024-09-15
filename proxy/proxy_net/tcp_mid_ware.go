package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
)

type TcpBaseMidWare struct {
}

// PreHandler Handler 前的 hook 方法
func (midWare *TcpBaseMidWare) PreHandler(iproxy_net.ITcpReq) {
}

// MsgHandler 处理拆包得到的 tcp 消息
func (midWare *TcpBaseMidWare) MsgHandler(iproxy_net.ITcpReq) {
}

// PostHandler Handler 后的 hook 方法
func (midWare *TcpBaseMidWare) PostHandler(iproxy_net.ITcpReq) {
}
