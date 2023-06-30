#include <iostream>
#include <unordered_map>
#include <string>

#include "../data/peer.h"

#ifndef MULTICAST_H
#define MULTICAST_H

void send_multicast_broadcast(int socket);
void create_multicast_socket();
void start_multicast_handler();
int get_multicast_socket();

std::unordered_map<struct Peer, std::string> get_peers(); //peer and name
void add_peer(struct Peer *peer, std::string name);
void remove_peer(struct Peer *peer);


#endif
