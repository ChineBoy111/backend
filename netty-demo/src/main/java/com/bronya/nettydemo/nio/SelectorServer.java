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
public class SelectorServer {

    public static void split(ByteBuffer buf) {
        buf.flip(); // buf 切换到读模式
        for (int i = 0; i < buf.limit(); i++) {
            if (buf.get(i) == '\n') { // 找到一个完整的数据包，读指针 position 不移动
                var len = i + 1 - buf.position(); // 数据包长度
                var tmpBuf = ByteBuffer.allocate(len);
                for (int j = 0; j < len; j++) {
                    var b = buf.get(); // 从 buf 中读出一个字节，读指针 position 向右移动一位
                    tmpBuf.put(b);  // 向 tmpBuf 中写入一个字节
                }
                tmpBuf.flip(); // tmpBuf 切换到读模式
                System.out.println("[DEBUG] " + Charset.forName("UTF-8").decode(buf).toString());
            } // 未找到一个完整的数据包，读指针 = 读取限制 (position == limit)
        }
        buf.compact(); // buf 紧凑
    }

    public static void main(String[] args) throws IOException {
        //* 创建 listener 套接字
        var listener = ServerSocketChannel.open();
        //* 非阻塞 listener
        listener.configureBlocking(false);
        //* 创建 selector
        var selector = Selector.open();
        //* 向 selector 中注册 listener
        var listenerKey = listener.register(selector, 0/* options 0 表示 selector 不关注任何事件 */, null/* attachments 附件 */);
        //* listener 关注 accept 事件
        listenerKey.interestOps(SelectionKey.OP_ACCEPT);
        //* listener 绑定本机 3333 端口，监听客户端的连接请求
        listener.bind(new InetSocketAddress(3333/* port */));
        System.out.println("SelectorServer listening on port 3333");
        while (true) {
            //* select 方法
            //* - 无关注事件发生：线程放弃 cpu，阻塞等待
            //* - 有关注事件发生，未处理/未取消：线程占用 cpu，处理/取消事件
            selector.select(); // 选择一个发生的关注事件
            //* keySet 是待处理/待取消事件 key 的集合
            var keySet = selector.selectedKeys();
            var iter = keySet.iterator();
            //* 处理/取消事件
            while (iter.hasNext()) {
                var key = iter.next();
                //! 处理/取消事件后，需要删除 keySet 中该事件的 key
                iter.remove();
                if (key.isAcceptable()) { //* 是 accept 事件
                    //* 获取 accept 事件对应的 listener
                    var listener_ = (ServerSocketChannel) key.channel(); // listener_ == listener
                    //* 非阻塞等待客户端的连接请求
                    var clientSocket = listener_.accept();
                    //! 非阻塞 clientSocket
                    //! 如果 clientSocket 中没有数据，read 方法返回 0，线程继续运行
                    clientSocket.configureBlocking(false);
                    var buf = ByteBuffer.allocate(16);
                    //* 向 selector 中注册 clientSocket，buf 作为 clientSocket 的附件
                    var clientSocketKey = clientSocket.register(selector, 0/* options */, buf/* attachments */);
                    //* clientSocket 关注 readable 可读事件
                    clientSocketKey.interestOps(SelectionKey.OP_READ);
                } else if (key.isReadable()) { //* 是 readable 事件
                    try {
                        //* 获取 readable 事件对应的 clientSocket
                        var clientSocket = (SocketChannel) key.channel();
                        //* 获取 clientSocket 的附件 buf
                        var buf = (ByteBuffer) key.attachment();
                        //* 客户端正常断开，read 方法返回 -1
                        var readBytes = clientSocket.read(buf);
                        if (readBytes == -1) {
                            key.cancel(); //* 取消事件
                        } else {
                            split(buf);
                            if (buf.position() == buf.limit()) { // 未找到一个完整的数据包
                                var newBuf = ByteBuffer.allocate(buf.capacity() * 2); // 扩容
                                buf.flip(); // buf 切换到读模式
                                newBuf.put(buf); // 将 buf 的数据拷贝到 newBuf
                                key.attach(newBuf); //* 更新 clientSocket 的附件
                            }
                        }
                    } catch (SocketException ignored) {
                        //* 客户端异常断开
                        key.cancel(); //* 取消事件
                    }
                }
            }
        }
    }
}
