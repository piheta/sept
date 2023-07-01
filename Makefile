CC = g++
CFLAGS = -Wall -Wextra -g
CXXFLAGS = $(CFLAGS)

SRCS = data/peer.c network/multicast.cc network/p2p.c node.cc
OBJS = $(SRCS:.c=.o)
TARGET = p2p

.PHONY: all clean

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) $(CXXFLAGS) $^ -o $@

%.o: %.c
	$(CC) $(CFLAGS) -x c++ -c $< -o $@

clean:
	rm -f */*.o
	rm -f node.o
	rm -rf p2p.dSYM
