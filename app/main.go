package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	node := RunNode()
	log.Printf("Hello World, my second hosts ID is %s\n", node.ID())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	// shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}
}

func RunNode() host.Host {
	ctx := context.Background()
	node, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/63785", "/ip4/0.0.0.0/tcp/63786/ws"),
		libp2p.Ping(false),
	)
	if err != nil {
		panic(err)
	}
	defer node.Close()
	log.Printf("Hello World, my hosts ID is %s\n", node.ID())

	return node
}
