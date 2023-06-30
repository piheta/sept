#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#ifndef PEER_H
#define PEER_H

struct Peer {
    const char* ip;
    int port;
};

struct sockaddr_in get_peer_from_multicast_msg();

struct Peer peer_constructor(const char* peer_ip, int peer_port);
void peer_destructor(struct Peer *peer);


#endif
