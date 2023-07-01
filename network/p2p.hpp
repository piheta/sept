#include <iostream>
#include <unordered_map>
#include <string>

#include "../data/peer.h"

#ifndef P2P_H
#define P2P_H

void create_p2p_socket(int sock_fd, int local_port, int remote_port, std::string remote_host);
void p2p_listen(int sock_fd);
void p2p_send(int sock_fd);


#endif
