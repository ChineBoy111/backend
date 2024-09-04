package com.bronya.nettydemo.nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;
import java.nio.charset.Charset;
import java.util.ArrayList;

//! 阻塞 IO
public class BlockServer {

    public static void main(String[] args) throws IOException {
        //* 创建 listener 套接字
        var listener = ServerSocketChannel.open();
        //* listener 绑定本机 3333 端口，监听客户端的连接请求
        listener.bind(new InetSocketAddress(3333/* port */));
        System.out.println("BlockServer listening on port 3333");
        var clientSocketList = new ArrayList<SocketChannel>();
        var buf = ByteBuffer.allocate(16);
        while (true) {
            //* 阻塞等待客户端的连接请求
            var clientSocket = listener.accept();
            //! 阻塞 listener
            //! 如果未收到客户端的连接请求，线程阻塞在 accept，不能 read
            clientSocketList.add(clientSocket);
            for (var socket : clientSocketList) {
                //* 阻塞等待从 socket 中读出数据，写入 buf（写模式）
                socket.read(buf);
                //! 阻塞 socket
                //! 如果 socket 中没有数据，线程阻塞在 read，不能 accept
                buf.flip(); //* buf 切换到读模式
                System.out.println("[DEBUG] " + Charset.forName("UTF-8").decode(buf).toString());
                buf.clear(); //* 清空 buf
            }
        }
    }
}
