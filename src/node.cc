#include <iostream>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <errno.h>
#include <netinet/in.h>
#include <poll.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <chrono>

#include "network/multicast.hh"
#include "network/p2p.h"

void start_node_multiplexing(int sock_fd, int m_sock_fd) {
    int ret;
    struct pollfd fds[3];

    fds[0].fd = 0; //stdin
    fds[1].fd = sock_fd;
    fds[2].fd = m_sock_fd;
    fds[0].events = POLLIN | POLLPRI;
    fds[1].events = POLLIN | POLLPRI;
    fds[2].events = POLLIN | POLLPRI;

    send_multicast_broadcast(m_sock_fd, seek);
    std::chrono::steady_clock::time_point startup_time = std::chrono::steady_clock::now();

    while (1) {
        ret = poll(fds, 3, -1);

        if (ret < 0) {
            printf("Error - poll returned error: %s\n", strerror(errno));
            break;
        }
        if (ret > 0) {

            /* poll() can always return these */
            if (fds[0].revents & (POLLERR | POLLHUP | POLLNVAL)) {
                printf("Error - poll indicated stdin error\n");
                break;
            }
            if (fds[1].revents & (POLLERR | POLLHUP | POLLNVAL)) {
                printf("Error - poll indicated socket error\n");
                break;
            }
            if (fds[2].revents & (POLLERR | POLLHUP | POLLNVAL)) {
                printf("Error - poll indicated m_socket error\n");
                break;
            }

            /* stdin */
            if (fds[0].revents & (POLLIN | POLLPRI)) {
                p2p_send(sock_fd);
            }

            /* p2p socket */
            if (fds[1].revents & (POLLIN | POLLPRI)) {
                p2p_listen(sock_fd);
            }

            /* multicast socket */
            if (fds[2].revents & (POLLIN | POLLPRI)) {
                multicast_handler(m_sock_fd, startup_time);
            }
        }
    }
}


int main() {
    std::string version = "v0.2.9";
    int sock_port = 50012;
    int m_sock_port = 50010;
    const char* m_sock_addr = "239.50.0.10";

    std::cout << "\e[1;1H\e[2J";
    std::cout << "\x1b[38;5;209m╔═╗╔═╗╔═╗╔╦╗ \x1b[0m│ " << version << std::endl;
    std::cout << "\x1b[38;5;214m╚═╗╠═ ╠═╝ ║  \x1b[0m│ :" << m_sock_port << std::endl;
    std::cout << "\x1b[38;5;215m╠═╝╚═╝╩   ╩  \x1b[0m│ :" << sock_port << std::endl;
    std::cout << "\x1b[38;5;220m~~~~~~~~~~~~~~~~~~~~~\x1b[0m" << std::endl;

    int sock_fd  = socket(AF_INET, SOCK_DGRAM, 0);
    int m_sock_fd = socket(AF_INET, SOCK_DGRAM, 0);
    create_p2p_socket(sock_fd, sock_port, sock_port, "0.0.0.0");
    create_multicast_socket(m_sock_fd, m_sock_port, m_sock_addr);
    start_node_multiplexing(sock_fd, m_sock_fd);

    send_multicast_broadcast(m_sock_fd, rm);
    close(sock_fd);
    close(m_sock_fd);
    return 0;
}
