package test

import (
	"bronya.com/proxy/proxy_net"
	"bronya.com/proxy/utils"
	"io"
	"log"
	"net"
	"sync"
	"testing"
)

func TestGlobal(t *testing.T) {
	log.Println(utils.Global)
}

func TestTcpPacKit(t *testing.T) {
	waitGroup := sync.WaitGroup{}
	defer waitGroup.Wait()

	listener, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Println("Server listen err", err.Error())
		return
	}

	// 服务器
	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()

		conn, err := listener.Accept()
		if err != nil {
			log.Println("Server accept err", err.Error())
			return
		}
		pacKit := proxy_net.NewTcpPacKit()
		go func(conn net.Conn) {
			waitGroup.Add(1)
			defer waitGroup.Done()

			for {
				// 第 1 次从 conn 中读，读出 pacHead (msgLen + msgId)
				pacHead := make([]byte, 8)
				readBytes, err := io.ReadFull(conn, pacHead)
				if err != nil || readBytes != 8 {
					log.Println("Read full err", err.Error())
					return
				}
				iMsg, err := pacKit.Unpack(pacHead)
				if err != nil {
					log.Println("Unpack err", err.Error())
					return
				}

				msg := iMsg.(*proxy_net.TcpMsg) // 类型断言
				if msg.GetLen() > 0 {           // 数据包头部 pacHead 的 msgLen > 0
					// 第 2 次从 conn 中读，读出 body (Data)
					msg.Data = make([]byte, msg.Len)
					readBytes, err = io.ReadFull(conn, msg.Data)
					if err != nil || uint32(readBytes) != msg.Len {
						log.Println("Read full err", err.Error())
						return
					}
				}
				log.Printf("==> Read msg.Len=%v, msg.Id=%v, msg.Data=%v\n", msg.Len, msg.Id, string(msg.Data))
			}
		}(conn)
	}()

	// 客户端
	conn, err := net.Dial("tcp4", "127.0.0.1:8080")
	//! 在闭包内封装错误处理
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Close err", err.Error())
		}
	}(conn)
	if err != nil {
		log.Println("Dial err", err.Error())
		return
	}
	pacKit := proxy_net.NewTcpPacKit()

	// 封装第 1 个 tcp 数据包 pac1
	msg1 := &proxy_net.TcpMsg{
		Len:  3,
		Id:   0,
		Data: []byte{'W', 'A', 'N'},
	}
	pac1, err := pacKit.Pack(msg1)
	if err != nil {
		log.Println("Pack err", err.Error())
		return
	}

	// 封装第 2 个 tcp 数据包 pac2
	msg2 := &proxy_net.TcpMsg{
		Len:  5,
		Id:   1,
		Data: []byte{'P', 'r', 'o', 'x', 'y'},
	}
	pac2, err := pacKit.Pack(msg2)
	if err != nil {
		log.Println("Pack err", err.Error())
		return
	}

	// tcp 粘包
	pac1 = append(pac1, pac2...)
	_ /* writeBytes */, err = conn.Write(pac1)
	if err != nil {
		log.Println("Write err", err.Error())
		return
	}
}
