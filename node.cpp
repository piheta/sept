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

#include "network/multicast.hpp"

void start_chat(int sock_fd, int m_sock_fd, struct sockaddr_in *peer) {
    int ret;
    ssize_t bytes;
    char input_buffer[1024];
    char output_buffer[1024];

    struct pollfd fds[3];

    fds[0].fd = 0; //stdin
    fds[1].fd = sock_fd;
    fds[2].fd = m_sock_fd;
    fds[0].events = POLLIN | POLLPRI;
    fds[1].events = POLLIN | POLLPRI;
    fds[2].events = POLLIN | POLLPRI;

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
                bytes = read(0, output_buffer, sizeof(output_buffer));
                if (bytes < 0) {
                    printf("Error - stdin error: %s\n", strerror(errno));
                    break;
                }
                if (strcmp(output_buffer, "exit\n") == 0) {
                    break;
                }
                bytes = sendto(sock_fd, output_buffer, bytes, 0, (struct sockaddr *)peer, sizeof(struct sockaddr_in));
                if (bytes < 0) {
                    printf("Error - sendto error: %s\n", strerror(errno));
                    break;
                }
                memset(output_buffer, 0, sizeof(output_buffer));
            }

            /* p2p socket */
            if (fds[1].revents & (POLLIN | POLLPRI)) {
                struct sockaddr_in peer_addr;
                socklen_t addrlen = sizeof(peer_addr);
                bytes = recvfrom(sock_fd, input_buffer, sizeof(input_buffer), 0, (struct sockaddr *)&peer_addr, &addrlen);
                if (bytes < 0) {
                    printf("Error - recvfrom error: %s\n", strerror(errno));
                    break;
                }
                if (bytes > 0) {
                    std::cout << inet_ntoa(peer_addr.sin_addr) << ":" << ntohs(peer_addr.sin_port) << ": " << input_buffer;
                }
                memset(input_buffer, 0, sizeof(input_buffer));
            }

            /* multicast socket */
            if (fds[2].revents & (POLLIN | POLLPRI)) {
                //start_multicast_handler();
            }
        }
    }
}


int main(int argc, char *argv[]) {
    unsigned long local_port;
    std::string remote_host;
    unsigned long remote_port;
    int sock_fd;
    struct sockaddr_in client_addr;
    struct sockaddr_in peer_addr;

    /* Parse command line arguments for port numbers */
    if (argc < 4) {
        printf("Usage: %s <local port> <remote host> <remote port>\n", argv[0]);
        return 1;
    }
    local_port = strtoul(argv[1], NULL, 0);
    remote_host = argv[2];
    remote_port = strtoul(argv[3], NULL, 0);


    /* Parse command line argument for remote host address */
    peer_addr.sin_family = AF_INET;
    peer_addr.sin_addr.s_addr = inet_addr(remote_host.c_str());
    peer_addr.sin_port = htons(remote_port);

    /* Create UDP socket */
    sock_fd = socket(AF_INET, SOCK_DGRAM, 0);
    if (sock_fd < 0) {
        printf("Error - failed to open socket: %s\n", strerror(errno));
        return 1;
    }

    /* Bind socket */
    client_addr.sin_family = AF_INET;
    client_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    client_addr.sin_port = htons(local_port);
    if (bind(sock_fd, (struct sockaddr *)(&client_addr), sizeof(client_addr)) < 0) {
        printf("Error - failed to bind socket: %s\n", strerror(errno));
        return 1;
    }

    start_chat(sock_fd, get_multicast_socket(), &peer_addr);

    close(sock_fd);
    return 0;
}
