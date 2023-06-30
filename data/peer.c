#include "peer.h"

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <arpa/inet.h>

struct Peer peer_constructor(const char *peer_ip, int peer_port) {
    struct Peer peer;
    peer.ip = peer_ip;
    peer.port = peer_port;

    return peer;
}

void peer_destructor(struct Peer *peer){
    free(peer);
}


struct sockaddr_in get_addr_from_multicast_msg() {
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;

    addr.sin_addr.s_addr = inet_addr("239.50.0.10");
    addr.sin_port = htons(50010);

    return addr;
}
