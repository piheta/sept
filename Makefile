CC = g++
CFLAGS = -Wall -Wextra -g

SRCS = data/peer.c network/multicast.cpp node.cpp
OBJS = $(SRCS:.c=.o)
TARGET = p2p

.PHONY: all clean

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) $^ -o $@

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -f data/peer.o
	rm -rf p2p.dSYM
