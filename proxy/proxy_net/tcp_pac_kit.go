package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"bronya.com/proxy/utils"
	"bytes"
	"encoding/binary"
	"errors"
)

type TcpPacKit struct {
}

func NewTcpPacKit() *TcpPacKit {
	return &TcpPacKit{}
}

// GetHeadLen 获取 tcp 数据包的 head 长度
func (pacKit *TcpPacKit) GetHeadLen() uint32 {
	return 8 // 4 bytes (Id uint32) + 4 bytes (Len uint32) = 8 bytes
}

// Pack tcp 封包，msg 结构体实例序列化为 packet 字节数组
func (pacKit *TcpPacKit) Pack(msg iproxy_net.ITcpMsg) ([]byte, error) {
	buf /* writer */ := bytes.NewBuffer([]byte{})
	// 向 buf 中 写入 Len
	if err := binary.Write(buf, binary.LittleEndian, msg.GetLen()); err != nil {
		return nil, err
	}
	// 向 buf 中 写入 Id
	if err := binary.Write(buf, binary.LittleEndian, msg.GetId()); err != nil {
		return nil, err
	}
	// 向 buf 中写入 Data
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	packet := buf.Bytes()
	return packet, nil
}

// Unpack tcp 拆包，packet 字节数组反序列化为 msg 结构体实例
func (pacKit *TcpPacKit) Unpack(packet []byte) (iproxy_net.ITcpMsg, error) {
	reader := bytes.NewReader(packet)
	msg := &TcpMsg{}
	// 从 byteArr 中读出 Len 到 msg
	if err := binary.Read(reader, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}
	// 从 byteArr 中读出 Id 到 msg
	if err := binary.Read(reader, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if msg.Len > utils.Global.TcpMaxPacketSize {
		return nil, errors.New("packet too big")
	}
	return msg, nil
}
