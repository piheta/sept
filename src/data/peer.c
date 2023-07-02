#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <arpa/inet.h>

struct Peer {
    const char* ip;
    int port;
};

struct Peer peer_constructor(const char *peer_ip, int peer_port) {
    struct Peer peer;
    peer.ip = peer_ip;
    peer.port = peer_port;

    return peer;
}

void peer_destructor(struct Peer *peer){
    free(peer);
}
