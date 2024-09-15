package proxy

import (
	"bronya.com/net-proxy/iproxy"
	"bronya.com/net-proxy/utils"
	"bytes"
	"encoding/binary"
	"errors"
)

type TcpPacKit struct {
}

var pacKit *TcpPacKit

// ! init 函数只执行 1 次
func init() {
	pacKit = &TcpPacKit{}
}

// NewTcpPacKit 创建 TcpPacKit 结构体变量
func NewTcpPacKit() *TcpPacKit {
	return pacKit
}

// GetHeadLen 获取 tcp 数据包的 head 长度
func (pacKit *TcpPacKit) GetHeadLen() uint32 {
	return 8 // 4 bytes (Id uint32) + 4 bytes (Len uint32) = 8 bytes
}

// Pack 封包，将 msg 结构体变量序列化为 packet 字节数组（tcp 消息 -> tcp 数据包）
func (pacKit *TcpPacKit) Pack(msg iproxy.ITcpMsg) ([]byte, error) {
	buf /* writer */ := bytes.NewBuffer([]byte{})
	// 向 buf 中 写入 msgLen
	if err := binary.Write(buf, binary.LittleEndian, msg.GetLen()); err != nil {
		return nil, err
	}
	// 向 buf 中 写入 msgId
	if err := binary.Write(buf, binary.LittleEndian, msg.GetId()); err != nil {
		return nil, err
	}
	// 向 buf 中写入 msgData
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	pac := buf.Bytes()
	return pac, nil
}

// Unpack 拆包，将 packet 字节数组反序列化为 msg 结构体变量（tcp 数据包 -> tcp 消息）
func (pacKit *TcpPacKit) Unpack(pac []byte) (iproxy.ITcpMsg, error) {
	reader := bytes.NewReader(pac)
	msg := &TcpMsg{}
	// 从 byteArr 中读出 msgLen 到 Msg.Len
	if err := binary.Read(reader, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}
	// 从 byteArr 中读出 msgId 到 Msg.Id
	if err := binary.Read(reader, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if msg.Len > utils.Global.TcpMaxPacSize {
		return nil, errors.New("tcp packet too big")
	}
	return msg, nil
}
