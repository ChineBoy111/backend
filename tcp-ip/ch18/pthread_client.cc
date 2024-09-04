//
// Created by Tiancheng on 2024/9/1.
//

#include <arpa/inet.h>
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <pthread.h>
#include <sys/socket.h>
#include <unistd.h>

#define BUF_SIZE 30
using any = void *;

any sendMsg(any arg) {
    int clientSocketFd = *(int *)arg;
    char buf[BUF_SIZE];
    while (true) {
        fgets(buf, BUF_SIZE - 10, stdin);
        if (strcasecmp(buf, "q\n") == 0) {
            close(clientSocketFd);
            return NULL;
        }

        //! 将 '\n' 替换为 '\0'
        buf[strlen(buf) - 1] = '\0';
        //! strlen(buf) 不计算 '\0'
        write(clientSocketFd, buf, strlen(buf) + 1);
    }
}

any recvMsg(any arg) {
    int clientSocketFd = *(int *)arg;
    char buf[BUF_SIZE];
    while (true) {
        int readLen = read(clientSocketFd, buf, BUF_SIZE);
        if (readLen <= 0) {
            return NULL;
        }
        printf("[INFO] Echo from server %s\n", buf);
    }
}

int main(int argc, char *argv[]) {
    if (argc != 3) {
        printf("Usage: %s <serverAddr> <serverPort>\n", argv[0]);
        exit(1);
    }

    int clientSocketFd = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);

    sockaddr_in serverAddr{};
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = inet_addr(argv[1]);
    serverAddr.sin_port = htons(atoi(argv[2]));

    if (connect(clientSocketFd, (sockaddr *)&serverAddr, sizeof(serverAddr)) ==
        -1) {
        printf("Error connected to server\n");
        exit(1);
    }

    pthread_t sendThreadId;
    pthread_t recvThreadId;

    pthread_create(&sendThreadId, NULL, sendMsg,
                   &clientSocketFd); // 创建发送线程
    pthread_create(&recvThreadId, NULL, recvMsg,
                   &clientSocketFd); // 创建接收线程

    pthread_join(sendThreadId, NULL); // 主进程阻塞等待发送进程终止
    pthread_join(recvThreadId, NULL); // 主进程阻塞接收发送进程终止
    close(clientSocketFd);
    return 0;
}