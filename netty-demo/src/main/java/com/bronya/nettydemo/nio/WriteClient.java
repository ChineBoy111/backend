package com.bronya.nettydemo.nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SocketChannel;

public class WriteClient {

    public static void main(String[] args) throws IOException {
        var clientSocket = SocketChannel.open();
        clientSocket.connect(new InetSocketAddress("172.25.154.2", 3333));
        int totalBytes = 0;
        while (true) {
            var buf = ByteBuffer.allocate(1024 * 1024);
            var readBytes = clientSocket.read(buf);
            totalBytes += readBytes;
            System.out.println("totalBytes = " + totalBytes);
            buf.clear();
        }
    }
}
