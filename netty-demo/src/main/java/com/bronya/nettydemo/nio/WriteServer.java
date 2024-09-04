package com.bronya.nettydemo.nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.net.SocketException;
import java.nio.ByteBuffer;
import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;
import java.nio.charset.Charset;

//! 非阻塞 IO
public class WriteServer {

    public static void main(String[] args) throws IOException {
        //* 创建 listener 套接字
        var listener = ServerSocketChannel.open();
        //* 非阻塞 listener
        listener.configureBlocking(false);
        //* listener 绑定本机 3333 端口，监听客户端的连接请求
        listener.bind(new InetSocketAddress(3333));
        //* 创建 selector
        var selector = Selector.open();
        //* 向 selector 中注册 listener，listener 关注 accept 事件
        listener.register(selector, SelectionKey.OP_ACCEPT);
        System.out.println("WriteServer listening on port 3333");
        while (true) {
            //* select 方法
            //* - 无关注事件发生：线程放弃 cpu，阻塞等待
            //* - 有关注事件发生，未处理/未取消：线程占用 cpu，处理/取消事件
            selector.select(); // 选择一个发生的关注事件
            //* keySet 是待处理事件 key 的集合
            var keySet = selector.selectedKeys();
            var iter = keySet.iterator();
            //* 处理/取消事件
            while (iter.hasNext()) {
                var key = iter.next();
                //! 处理/取消事件后，需要删除 keySet 中该事件的 key
                iter.remove();
                if (key.isAcceptable()) { //* 是 accept 事件
                    //* 非阻塞等待客户端的连接请求
                    var clientSocket = listener.accept();
                    //! 非阻塞 clientSocket
                    //! 如果 clientSocket 中没有数据，read 方法返回 0，线程继续运行
                    clientSocket.configureBlocking(false);
                    //* 向 selector 中注册 clientSocket，buf 作为 clientSocket 的附件
                    var clientSocketKey = clientSocket.register(selector, 0, null);
                    var builder = new StringBuilder();
                    for (int i = 0; i < 3_000_000; i++) {
                        builder.append("a");
                    }
                    var buf = Charset.defaultCharset().encode(builder.toString());
                    //! 第 1 次写
                    //* 将 buf 中的数据写入 clientSocket
                    var writeBytes /* 写入 clientSocket 的字节数 */ = clientSocket.write(buf);
                    System.out.println("[Accept] writeBytes = " + writeBytes);
                    if (buf.hasRemaining()) { //? buf 中有剩余数据
                        //* clientSocket 关注 writable 可写事件
                        clientSocketKey.interestOps(SelectionKey.OP_WRITE);
                        //* buf 作为 clientSocketKey 的附件
                        clientSocketKey.attach(buf);
                    } //? buf 中没有剩余数据，触发 writable 可写事件
                } else if (key.isWritable()) { //* 是 writable 可写事件
                    try {
                        //* 获取 writable 事件对应的 clientSocket
                        var clientSocket = (SocketChannel) key.channel();
                        //* 获取 clientSocket 的附件 buf
                        var buf = (ByteBuffer) key.attachment();
                        //! 第 2~n 次写
                        //* 将 buf 中的数据写入 clientSocket
                        var writeBytes /* 写入 clientSocket 的字节数 */ = clientSocket.write(buf);
                        System.out.println("[Writable] writeBytes = " + writeBytes);
                        if (!buf.hasRemaining()) { //? buf 中没有剩余数据
                            //* 注销 clientSocket 的附件 buf
                            key.attach(null);
                            //* clientSocket 不再关注 writable 可写事件
                            key.interestOps(0);
                        } //? buf 中有剩余数据
                    } catch (SocketException ignored) {
                        //* 客户端异常断开
                        key.cancel(); //* 取消事件
                    }
                }
            }
        }
    }
}
