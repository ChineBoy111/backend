package com.bronya.nettydemo.nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;
import java.nio.charset.Charset;
import java.util.ArrayList;

//! 非阻塞 IO
public class NonBlockServer {

    public static void main(String[] args) throws IOException {
        //* 创建 listener 套接字
        var listener = ServerSocketChannel.open();
        //* 非阻塞 listener
        listener.configureBlocking(false);
        //* listener 绑定本机 3333 端口，监听客户端的连接请求
        listener.bind(new InetSocketAddress(3333/* port */));
        System.out.println("NonBlockServer listening on port 3333");
        var socketList = new ArrayList<SocketChannel>();
        var buf = ByteBuffer.allocate(16);
        while (true) {
            //* 非阻塞等待客户端的连接请求
            var clientSocket = listener.accept();
            //! 非阻塞 listener
            //! 如果未收到客户端的连接请求，accept 方法返回 null，线程继续运行
            if (clientSocket != null) {
                //* 非阻塞 clientSocket
                clientSocket.configureBlocking(false);
                socketList.add(clientSocket);
            }

            for (SocketChannel socket : socketList) {
                //* 非阻塞等待从 socket 中读出数据，写入 buf（写模式）
                var readBytes = socket.read(buf);
                //! 非阻塞 socket
                //! 如果 socket 中没有数据，read 方法返回 0，线程继续运行 
                if (readBytes > 0) {
                    buf.flip(); //* buf 切换到读模式
                    System.out.println("[DEBUG] " + Charset.forName("UTF-8").decode(buf).toString());
                    buf.clear(); //* 清空 buf
                }
            }
        }
    }
}
