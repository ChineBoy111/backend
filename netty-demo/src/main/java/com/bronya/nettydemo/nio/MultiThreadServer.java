package com.bronya.nettydemo.nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.ClosedChannelException;
import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;
import java.nio.charset.Charset;
import java.util.concurrent.ConcurrentLinkedQueue;

public class MultiThreadServer {

    public static void main(String[] args) throws IOException {
        Thread.currentThread().setName("boss");
        var listener = ServerSocketChannel.open();
        listener.configureBlocking(false);
        var connSelector = Selector.open();
        var listenerKey = listener.register(connSelector, 0, null);
        listenerKey.interestOps(SelectionKey.OP_ACCEPT);
        listener.bind(new InetSocketAddress(3333));

        var worker0 = new Worker("worker0");
        while (true) {
            connSelector.select(); // 阻塞等待选择一个发生的关注事件
            var iter = connSelector.selectedKeys().iterator();
            while (iter.hasNext()) {
                var key = iter.next(); // key == listenerKey
                iter.remove();
                if (key.isAcceptable()) {
                    var listener_ = (ServerSocketChannel) key.channel(); // listener_ == listener
                    var clientSocket = listener_.accept();
                    clientSocket.configureBlocking(false);
                    System.out.println("Accepted connection from " + clientSocket.getRemoteAddress());
                    worker0.work(clientSocket);
                }
            }
        }
    }

    // 1. this.rwSelector = Selector.open();
    // 2. clientSocket.register(worker0.rwSelector, SelectionKey.OP_READ, null);
    // 3. rwSelector.select();
    static class Worker implements Runnable {

        private String name;
        private volatile boolean init = false;
        private Selector rwSelector;
        private Thread thread;
        private ConcurrentLinkedQueue<Runnable> taskQueue = new ConcurrentLinkedQueue<>();

        public Worker(String name) throws IOException {
            this.name = name;
        }

        public void work(SocketChannel clientSocket) throws IOException {
            if (!init) {
                this.rwSelector = Selector.open();
                //! 虚拟线程
                this.thread = Thread.ofVirtual().name(this.name).start(this);
                init = true;
            }
            taskQueue.add(() -> {
                try {
                    //* 向 worker0.rwSelector 中注册 clientSocket，clientSocket 关注 readable 可读事件
                    clientSocket.register(this.rwSelector, SelectionKey.OP_READ, null);
                } catch (ClosedChannelException e) {
                    e.printStackTrace();
                }
            });
            rwSelector.wakeup();
        }

        @Override
        public void run() {
            while (true) {
                try {
                    rwSelector.select(); // 阻塞等待选择一个发生的关注事件
                    var task = taskQueue.poll();
                    if (task != null) {
                        //* 向 worker0.rwSelector 中注册 clientSocket，clientSocket 关注 readable 可读事件
                        task.run();
                    }
                    var iter = rwSelector.selectedKeys().iterator();
                    while (iter.hasNext()) {
                        var key = iter.next();
                        iter.remove();
                        if (key.isReadable()) {
                            var buf = ByteBuffer.allocate(30);
                            var clientSocket = (SocketChannel) key.channel();
                            clientSocket.read(buf);
                            buf.flip();
                            System.out.println("[DEBUG] " + Charset.forName("UTF-8").decode(buf).toString());
                        } // key.isWritable()
                    }
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
