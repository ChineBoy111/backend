cmake_minimum_required(VERSION 3.21)

project(tcp_ip)

set(CMAKE_CXX_STANDARD 23)

set(CMAKE_GENERATOR "Ninja")

#! sudo apt install libboost-all-dev
# set(Boost_LIB_DIR "/usr/include/x86_64-linux-gnu")
# set(Boost_INCLUDE_DIR "/usr/include")

# include_directories("/* ... */")

# find_package(Boost COMPONENTS system filesystem program_options REQUIRED)
# target_link_libraries(main ${Boost_INCLUDE_DIR})

add_executable(ch01_server "ch01/server.cc")
add_executable(ch01_client "ch01/client.cc")
add_executable(ch01_test_fd "ch01/test_fd.cc")
add_executable(ch01_test_rw "ch01/test_rw.cc")

add_executable(ch02_tcp_server "ch02/tcp_server.cc")
add_executable(ch02_tcp_client "ch02/tcp_client.cc")

add_executable(ch03_test_inet "ch03/test_inet.cc")

add_executable(ch04_echo_server "ch04/echo_server.cc")
add_executable(ch04_echo_client "ch04/echo_client.cc")

add_executable(ch06_udp_server "ch06/udp_server.cc")
add_executable(ch06_udp_client "ch06/udp_client.cc")
add_executable(ch06_conn_client "ch06/conn_client.cc")

add_executable(ch07_file_server "ch07/file_server.cc")
add_executable(ch07_file_client "ch07/file_client.cc")

add_executable(ch08_test_gethostbyname "ch08/test_gethostbyname.cc")
add_executable(ch08_test_gethostbyaddr "ch08/test_gethostbyaddr.cc")

add_executable(ch09_test_sockopt "ch09/test_sockopt.cc")
add_executable(ch09_reuseaddr_server "ch09/reuseaddr_server.cc")

add_executable(ch10_test_fork "ch10/test_fork.cc")
add_executable(ch10_zombie "ch10/zombie.cc")
add_executable(ch10_kill_zombie "ch10/kill_zombie.cc")
add_executable(ch10_test_wait "ch10/test_wait.cc")
add_executable(ch10_test_waitpid "ch10/test_waitpid.cc")
add_executable(ch10_test_signal "ch10/test_signal.cc")
add_executable(ch10_test_sigaction "ch10/test_sigaction.cc")
add_executable(ch10_proc_server "ch10/proc_server.cc")
add_executable(ch10_proc_client "ch10/proc_client.cc")

add_executable(ch11_test_pipe1 "ch11/test_pipe1.cc")
add_executable(ch11_test_pipe2 "ch11/test_pipe2.cc")
add_executable(ch11_test_pipe3 "ch11/test_pipe3.cc")
add_executable(ch11_pipe_server "ch11/pipe_server.cc")

add_executable(ch12_test_select "ch12/test_select.cc")
add_executable(ch12_select_server "ch12/select_server.cc")

add_executable(ch13_oob_server "ch13/oob_server.cc")
add_executable(ch13_oob_client "ch13/oob_client.cc")
add_executable(ch13_peek_server "ch13/peek_server.cc")
add_executable(ch13_peek_client "ch13/peek_client.cc")
add_executable(ch13_test_writev "ch13/test_writev.cc")
add_executable(ch13_test_readv "ch13/test_readv.cc")

add_executable(ch14_multicast_server "ch14/multicast_server.cc")
add_executable(ch14_multicast_client "ch14/multicast_client.cc")
add_executable(ch14_broadcast_server "ch14/broadcast_server.cc")
add_executable(ch14_broadcast_client "ch14/broadcast_client.cc")

add_executable(ch15_io "ch15/io.cc")
add_executable(ch15_fd_fp "ch15/fd_fp.cc")
add_executable(ch15_echo_server "ch15/echo_server.cc")
add_executable(ch15_echo_client "ch15/echo_client.cc")

add_executable(ch16_server1 "ch16/server1.cc")
add_executable(ch16_server2 "ch16/server2.cc")
add_executable(ch16_client "ch16/client.cc")
add_executable(ch16_test_dup "ch16/test_dup.cc")

add_executable(ch17_epoll_server "ch17/epoll_server.cc")

add_executable(ch18_pthread1 "ch18/pthread1.cc")
add_executable(ch18_pthread2 "ch18/pthread2.cc")
add_executable(ch18_pthread3 "ch18/pthread3.cc")
add_executable(ch18_pthread4 "ch18/pthread4.cc")
add_executable(ch18_pthread_server "ch18/pthread_server.cc")
add_executable(ch18_pthread_client "ch18/pthread_client.cc")

add_executable(ch24_http_server "ch24/http_server.cc")

# 查找 pthread 库
find_package(Threads REQUIRED)

# 链接 pthread 库
target_link_libraries(ch18_pthread1 Threads::Threads)
target_link_libraries(ch18_pthread2 Threads::Threads)
target_link_libraries(ch18_pthread3 Threads::Threads)
target_link_libraries(ch18_pthread4 Threads::Threads)
target_link_libraries(ch18_pthread_server Threads::Threads)
target_link_libraries(ch18_pthread_client Threads::Threads)
target_link_libraries(ch24_http_server Threads::Threads)