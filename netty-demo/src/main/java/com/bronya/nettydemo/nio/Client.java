package com.bronya.nettydemo.nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.channels.SocketChannel;
import java.nio.charset.Charset;

public class Client {

    public static void main(String[] args) throws IOException {
        //* 创建 clientSocket 套接字
        var clientSocket = SocketChannel.open();
        System.out.println("Connecting to 127.0.0.1:3333");
        //* 向服务器发送连接请求
        clientSocket.connect(new InetSocketAddress("localhost", 3333));

        // clientSocket.write(Charset.defaultCharset().encode("0123\n456789abcdef"));
        // clientSocket.write(Charset.defaultCharset().encode("0123456789abcdefXXXX\n"));
        // clientSocket.close();//* 断开 socket 连接

        clientSocket.write(Charset.defaultCharset().encode("0123456789abcdef"));
        System.in.read();
    }
}
