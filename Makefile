CC = g++
CFLAGS = -Wall -Wextra -g
CXXFLAGS = $(CFLAGS)

SRCS = src/data/peer.c src/network/multicast.cc src/network/p2p.c src/node.cc
OBJS = $(SRCS:.c=.o)
TARGET = sept

.PHONY: all clean

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) $(CXXFLAGS) $^ -o $@

%.o: %.c
	$(CC) $(CFLAGS) -x c++ -c $< -o $@

clean:
	rm -f */*.o
	rm -f */*/*.o
	rm -f node.o
	rm -rf sept.dSYM
	rm -f .DS_Store
	rm -f */.DS_Store
	rm -f */*/.DS_Store
