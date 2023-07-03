#include <sys/socket.h>
#include <arpa/inet.h>
#include <errno.h>
#include <netinet/in.h>
#include <poll.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "p2p.h"
#include "../data/peer.h"

ssize_t bytes;
struct sockaddr_in client_addr;
struct sockaddr_in peer_addr;
const char* remote_host = "0.0.0.0";

void create_p2p_socket(int sock_fd, int local_port, int remote_port, const char* remote_host) {
    peer_addr.sin_family = AF_INET;
    peer_addr.sin_addr.s_addr = inet_addr(remote_host);
    peer_addr.sin_port = htons(remote_port);

    /* Create UDP socket */
    if (sock_fd < 0) {
        printf("Error - failed to open socket: %s\n", strerror(errno));
    }

    /* Bind socket */
    client_addr.sin_family = AF_INET;
    client_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    client_addr.sin_port = htons(local_port);
    if (bind(sock_fd, (struct sockaddr *)(&client_addr), sizeof(client_addr)) < 0) {
        printf("Error - failed to bind socket: %s\n", strerror(errno));
        close(sock_fd);
    }
}

void p2p_listen(int sock_fd) {
    char input_buffer[1024];
    struct sockaddr_in peer_addr;
    socklen_t addrlen = sizeof(peer_addr);
    bytes = recvfrom(sock_fd, input_buffer, sizeof(input_buffer), 0, (struct sockaddr *)&peer_addr, &addrlen);
    if (bytes < 0) {
        printf("Error - recvfrom error: %s\n", strerror(errno));
    }
    if (bytes > 0) {
        printf("%s:%d: %s", inet_ntoa(peer_addr.sin_addr), ntohs(peer_addr.sin_port), input_buffer);
    }
    memset(input_buffer, 0, sizeof(input_buffer));
}

void command(char output_buffer[1024]) {
    if (strncmp(output_buffer, "!help", 5) == 0) {
        printf("\e[1m" "!help " "\e[m" "- this list \n");
        printf("\e[1m" "!list " "\e[m" "- list online nodes \n");
        printf("\e[1m" "!peer <id> " "\e[m" "- set the node to send messages to \n");
        printf("\e[1m" "!exit " "\e[m" "- exit the program \n");
    }

    else if (strncmp(output_buffer, "!exit", 5) == 0) {
        exit(0);
    }

    else if (strncmp(output_buffer, "!peer", 5) == 0) {
        char* hostStart = output_buffer + 6;  // Add 1 to skip the space after the prefix
        char* hostEnd = output_buffer + strlen(output_buffer);

        // Calculate the length of the host portion
        size_t hostLength = hostEnd - hostStart;
        char remote_host[hostLength + 1];  // +1 for null-terminator
        strncpy(remote_host, hostStart, hostLength);
        peer_addr.sin_addr.s_addr = inet_addr(remote_host);
        printf("\x1b[32m" "success!" "\x1b[0m" "\n");
        remote_host[hostLength] = '\0';
    }

    else if (strncmp(output_buffer, "!list", 5) == 0) {
        //printf("10.0.0.1 \n");
    }

    else {
        printf("\x1b[38;5;9m" "error" "\x1b[0m" " - command not found, try " "\e[1m" "!help" "\e[m" " \n");
    }
}

void p2p_send(int sock_fd) {
    char output_buffer[1024];
    memset(output_buffer, 0, sizeof(output_buffer));
    bytes = read(0, output_buffer, sizeof(output_buffer));
    if (bytes < 0) {
        printf("Error - stdin error: %s\n", strerror(errno));
        return;
    }

    if (strncmp(output_buffer, "!", 1) == 0) {
        command(output_buffer);
        return;
    }

    bytes = sendto(sock_fd, output_buffer, bytes, 0, (struct sockaddr *)&peer_addr, sizeof(struct sockaddr_in));
    if (bytes < 0) {
        printf("Error - sendto error: %s\n", strerror(errno));
        return;
    }
}
