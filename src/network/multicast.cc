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

#include "multicast.hh"
#include "../data/peer.h"

struct sockaddr_in multicast_addr;
struct sockaddr_in groupSock;
struct ip_mreq m_group;
char m_databuf[1024];

std::unordered_map<std::string,struct Peer> known_hosts;

void send_multicast_broadcast(int m_sock_fd, Broadcast broadcast_type) {
    struct in_addr localInterface;

    groupSock.sin_family = AF_INET;
    groupSock.sin_addr.s_addr = inet_addr("239.50.0.10");
    groupSock.sin_port = htons(50010);
    localInterface.s_addr = htonl(INADDR_ANY);
    if (setsockopt(m_sock_fd, IPPROTO_IP, IP_MULTICAST_IF, (char *)&localInterface, sizeof(localInterface)) < 0) {
        perror("Setting local interface error");
        return;
    }

    char databuf[1024];
    int datalen = sizeof(databuf);
    switch (broadcast_type) {
        case seek:
            strcpy(databuf,"sept,0,uuid");
            break;
        case reply:
            strcpy(databuf,"sept,1,uuid");
            break;
        case rm:
            strcpy(databuf,"sept,2,uuid");
            break;
    }

    if (sendto(m_sock_fd, databuf, datalen, 0, (struct sockaddr*)&groupSock, sizeof(groupSock)) < 0) {
        perror("Sending datagram message error");
        return;
    }
}


void create_multicast_socket(int m_sock_fd, int port, const char* ip){
    if (m_sock_fd < 0) {
        perror("Opening datagram socket error");
        return;
    }

    int reuse = 1;
    if (setsockopt(m_sock_fd, SOL_SOCKET, SO_REUSEADDR, (char *)&reuse, sizeof(reuse)) < 0) {
        perror("Setting SO_REUSEADDR error");
        close(m_sock_fd);
        return;
    }

    multicast_addr.sin_family = AF_INET;
    multicast_addr.sin_addr.s_addr = INADDR_ANY;
    multicast_addr.sin_port = htons(port);
    if (bind(m_sock_fd, (struct sockaddr *)&multicast_addr, sizeof(multicast_addr))) {
        perror("Binding datagram socket error");
        close(m_sock_fd);
        return;
    }

    /* Join the multicast group on the local ip */
    m_group.imr_multiaddr.s_addr = inet_addr(ip);
    m_group.imr_interface.s_addr = htonl(INADDR_ANY);
    if (setsockopt(m_sock_fd, IPPROTO_IP, IP_ADD_MEMBERSHIP, (char *)&m_group, sizeof(m_group)) < 0) {
        perror("Adding multicast m_group error");
        close(m_sock_fd);
        return;
    }
}

void multicast_handler(int m_sock_fd) {
    struct sockaddr_in senderAddr;
    socklen_t addrLen = sizeof(senderAddr);
    if (recvfrom(m_sock_fd, m_databuf, sizeof(m_databuf), 0, (struct sockaddr*)&senderAddr, &addrLen) < 0) {
        perror("Receiving multigram datagram message error");
        return;
    }

    // stop if multicast traffic is not from sept
    if (strncmp(m_databuf, "sept", 4) != 0) {
        return;
    }

    switch (m_databuf[5]) {
        case seek:
            send_multicast_broadcast(m_sock_fd, reply);
            break;
        case reply:
            //add to known hosts
            struct Peer peer;
            peer = peer_constructor(inet_ntoa(senderAddr.sin_addr), ntohs(senderAddr.sin_port));
            known_hosts[inet_ntoa(senderAddr.sin_addr)] = peer;

            std::cout << known_hosts.size() << std::endl;

            break;
        case rm:
            //rm from known hosts
            break;
    }

    //printf("%s:%d: %s\n", inet_ntoa(senderAddr.sin_addr), ntohs(senderAddr.sin_port), m_databuf);

}
