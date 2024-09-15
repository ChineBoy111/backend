package proxy_net

import "bronya.com/proxy/iproxy_net"

type TcpReq struct {
	Conn iproxy_net.ITcpConn
	Msg  iproxy_net.ITcpMsg
}

// GetConn 获取 TcpConn 结构体变量
func (req *TcpReq) GetConn() iproxy_net.ITcpConn {
	return req.Conn
}

// GetMsgLen 获取 tcp 消息的长度
func (req *TcpReq) GetMsgLen() uint32 {
	return req.Msg.GetLen()
}

// GetMsgId 获取 tcp 消息的 id
func (req *TcpReq) GetMsgId() uint32 {
	return req.Msg.GetId()
}

// GetMsgData 获取 tcp 消息的数据
func (req *TcpReq) GetMsgData() []byte {
	return req.Msg.GetData()
}
