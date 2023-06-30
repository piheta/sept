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

#include "multicast.hpp"
#include "../data/peer.h"

int m_sock_fd;
struct sockaddr_in multicast_addr;
struct ip_mreq m_group;
struct sockaddr_in groupSock;
char m_databuf[1024];

void send_multicast_broadcast(int socket) {
    struct sockaddr_in groupSock;
    struct in_addr localInterface;

    char databuf[1024] = "seek";
    int datalen = sizeof(databuf);
    groupSock.sin_family = AF_INET;
    groupSock.sin_addr.s_addr = inet_addr("239.50.0.10");
    groupSock.sin_port = htons(50010);
    localInterface.s_addr = htonl(INADDR_ANY);
    if (setsockopt(socket, IPPROTO_IP, IP_MULTICAST_IF, (char *)&localInterface, sizeof(localInterface)) < 0) {
        perror("Setting local interface error");
    }

    if (sendto(socket, databuf, datalen, 0, (struct sockaddr*)&groupSock, sizeof(groupSock)) < 0) {
        perror("Sending datagram message error");
    }
}


void create_multicast_socket(){
    m_sock_fd = socket(AF_INET, SOCK_DGRAM, 0);
    if (m_sock_fd < 0) {
        perror("Opening datagram socket error");
    }

    int reuse = 1;
    if (setsockopt(m_sock_fd, SOL_SOCKET, SO_REUSEADDR, (char *)&reuse, sizeof(reuse)) < 0) {
        perror("Setting SO_REUSEADDR error");
        close(m_sock_fd);
    }


    memset((char *)&multicast_addr, 0, sizeof(multicast_addr));
    multicast_addr.sin_family = AF_INET;
    multicast_addr.sin_addr.s_addr = INADDR_ANY;
    multicast_addr.sin_port = htons(50010);
    if (bind(m_sock_fd, (struct sockaddr *)&multicast_addr, sizeof(multicast_addr))) {
        perror("Binding datagram socket error");
        close(m_sock_fd);
    }

    /* Join the multicast group on the local ip */
    m_group.imr_multiaddr.s_addr = inet_addr("239.50.0.10");
    m_group.imr_interface.s_addr = htonl(INADDR_ANY);
    if (setsockopt(m_sock_fd, IPPROTO_IP, IP_ADD_MEMBERSHIP, (char *)&m_group, sizeof(m_group)) < 0) {
        perror("Adding multicast m_group error");
        close(m_sock_fd);
    }
}

void start_multicast_handler(){

    //listen
    if (read(m_sock_fd, m_databuf, sizeof(m_databuf)) < 0) {
        perror("Reading multigram datagram message error");
    } else {
        printf("multicast: \"%s\"\n", m_databuf);
    }

    //add to peer set
    //reply
    //send_multicast_broadcast(m_sock_fd);
}

int get_multicast_socket(){
    return m_sock_fd;
}
