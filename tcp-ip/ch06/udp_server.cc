//
// Created by Tiancheng on 2024/8/18.
//

#include <arpa/inet.h>
#include <cstdio>
#include <cstdlib>
#include <sys/socket.h>
#include <unistd.h>

#define BUF_SIZE 32

int main(int argc, char **argv) {
    if (argc != 2) {
        printf("Usage: %s <port>\n", argv[0]);
        exit(1);
    }
    // 调用 socket 函数，创建 UDP socket 套接字
    int serverSocketFd = socket(PF_INET, SOCK_DGRAM, IPPROTO_UDP);
    if (serverSocketFd == -1) {
        perror("Error created socket");
    }
    sockaddr_in serverAddr{};
    serverAddr.sin_family = AF_INET; // IPv4 协议族
    // htonl 函数将一个 32 位（4 字节）的 int 整数从主机字节序转换为网络字节序
    serverAddr.sin_addr.s_addr =
        htonl(INADDR_ANY); // 0.0.0.0 接受所有 IP 地址的 TCP/UDP 连接
    // htons 函数将一个 16 位（2 字节）的 short 整数从主机字节序转换为网络字节序
    serverAddr.sin_port = htons(atoi(argv[1])); // 端口 = 第 1 个命令行参数
    //* 调用 bind 函数，给 socket 套接字分配 IP 地址和端口
    if (bind(serverSocketFd, (sockaddr *)&serverAddr, sizeof(serverAddr)) ==
        -1 /* 绑定 IP 地址、端口，成功时返回 0，失败时返回 -1 */) {
        perror("Error bound IP addr and port");
    }

    char buf[BUF_SIZE];
    sockaddr_in clientAddr{};
    socklen_t clientAddrLen = sizeof(clientAddr);

    while (true) {
        ssize_t recvBytes = recvfrom(serverSocketFd, buf, BUF_SIZE, 0,
                                   (sockaddr *)&clientAddr, &clientAddrLen);
        if (recvBytes == -1) {
            perror("[ERROR] Input is empty");
            break;
        }
        printf("buf = %s\n", buf);
        printf("clientAddr.sin_addr.s_addr = %d\n", clientAddr.sin_addr.s_addr);
        printf("clientAddr.sin_port = %d\n", clientAddr.sin_port);
        // if (clientAddr.sin_port == 0)
        //     break;
        sendto(serverSocketFd, buf, recvBytes, 0, (sockaddr *)&clientAddr,
               clientAddrLen);
    }
    close(serverSocketFd);
    return 0;
}
