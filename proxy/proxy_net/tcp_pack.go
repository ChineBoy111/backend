package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"bytes"
	"encoding/binary"
)

type TcpPack struct {
}

func NewTcpPack() *TcpPack {
	return &TcpPack{}
}

// GetHeadSize 获取 tcp 数据包的 head 头部大小
func (packer *TcpPack) GetHeadSize() uint32 {
	return 8 // 4 bytes (MsgId uint32) + 4 bytes (MsgLen uint32) = 8 bytes
}

// Pack tcp 封包
func (packer *TcpPack) Pack(msg iproxy_net.ITcpMsg) ([]byte, error) {
	buf /* writer */ := bytes.NewBuffer([]byte{})
	// 向 buf 中 写入 MsgLen
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// 向 buf 中 写入 MsgId
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 向 buf 中写入 MsgData
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	packet := buf.Bytes()
	return packet, nil
}

// Unpack tcp 拆包
func (packer *TcpPack) Unpack(packet []byte) (iproxy_net.ITcpMsg, error) {
	reader := bytes.NewReader(packet)
	msg := &TcpMsg{}
	// 从 packet 中读出 MsgLen
	if err := binary.Read(reader, binary.LittleEndian, &msg.MsgLen); err != nil {
		return nil, err
	}
	// 从 packet 中读出 MsgId
	if err := binary.Read(reader, binary.LittleEndian, &msg.MsgId); err != nil {
		return nil, err
	}
	// 从 packet 中读出 MsgData
	if err := binary.Read(reader, binary.LittleEndian, &msg.MsgData); err != nil {
		return nil, err
	}
	return msg, nil
}
