#include <iostream>
#include <unordered_map>
#include <string>

#include "../data/peer.h"

#ifndef MULTICAST_H
#define MULTICAST_H

enum Broadcast { seek = '0', reply = '1', rm = '2' }; // used in multicast messages, hence char

void send_multicast_broadcast(int m_sock_fd, Broadcast broadcast_type);
void create_multicast_socket(int m_sock_fd, int port, const char* ip);
void multicast_handler(int m_sock_fd, std::chrono::steady_clock::time_point startup_time);

std::unordered_map<struct Peer, std::string> get_peers(); //peer and name
void add_peer(struct Peer *peer, std::string name);
void remove_peer(struct Peer *peer);


#endif
