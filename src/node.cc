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
                multicast_handler(m_sock_fd);
            }
        }
    }
}


int main() {

    int sock_fd  = socket(AF_INET, SOCK_DGRAM, 0);
    int m_sock_fd = socket(AF_INET, SOCK_DGRAM, 0);

    printf("enter remote host ip:\n");
    char remote_host[50];
    scanf("%s", remote_host);

    create_multicast_socket(m_sock_fd, 50010, "239.50.0.10");
    create_p2p_socket(sock_fd, 70015, 70015, remote_host);
    start_node_multiplexing(sock_fd, m_sock_fd);

    send_multicast_broadcast(m_sock_fd, rm);
    close(sock_fd);
    close(m_sock_fd);
    return 0;
}
