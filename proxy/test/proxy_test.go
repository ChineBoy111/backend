package test

import (
	"bronya.com/proxy/proxy_net"
	"bronya.com/proxy/utils"
	"io"
	"log"
	"net"
	"testing"
)

func TestGlobal(t *testing.T) {
	log.Println(utils.Global)
}

func TestTcpPacket(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Println("Server listen err:", err.Error())
		return
	}

	// 服务器
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Server accept err:", err.Error())
				return
			}
			go func(conn net.Conn) {
				pacKit := proxy_net.NewTcpPacKit()
				// 第 1 次从 conn 中读，读出 head (Len + Id)
				head := make([]byte, 8)
				readBytes, err := io.ReadFull(conn, head)
				if err != nil || readBytes != 8 {
					log.Println("Read full err:", err.Error())
					return
				}
				iMsg, err := pacKit.Unpack(head)
				if err != nil {
					log.Println("Unpack err:", err.Error())
					return
				}

				msg := iMsg.(*proxy_net.TcpMsg) // 类型断言
				if msg.GetLen() > 0 {           // head 头部中 Len > 0
					// 第 2 次从 conn 中读，读出 body (Data)
					msg.Data = make([]byte, msg_.MsgLen)
					readBytes, err = io.ReadFull(conn, msg.Data)
					if err != nil || uint32(readBytes) != msg.Len {
						log.Println("Read full err:", err.Error())
						return
					}
				}
				log.Printf("Read Id = %v, Data = %v\n", msg.Id, msg.Data)
			}(conn)
		}
	}()

	// 客户端
	conn, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Println("Dial err:", err.Error())
		return
	}
	pacKit := proxy_net.NewTcpPacKit()
	// 粘包
	msg1 := &proxy_net.TcpMsg{
		Len:  5,
		Id:   0,
		Data: []byte{'W', 'a', 'n'},
	}
}
