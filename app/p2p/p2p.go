package p2p

import (
	"context"
	"fmt"
	"time"

	"github.com/blitzshare/blitzshare.bootstrap.node/app/dependencies"
	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	noise "github.com/libp2p/go-libp2p-noise"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	log "github.com/sirupsen/logrus"
)

func PrintNodeInfo(node host.Host) {
	log.Println()
	log.Printf(" - NODE ID: %s", node.ID())
	log.Printf(" - NODE ADDR: %v", node.Addrs())
	peers := node.Network().Peers()
	log.Infoln("Connected Peers", peers)
}

func PrintState(node host.Host) {
	for {
		PrintNodeInfo(node)
		time.Sleep(10 * time.Second)
	}
}

func RunNode(deps *dependencies.Dependencies) (host.Host, error) {
	ctx := context.Background()
	// Set your own keypair
	defaultConfigPriv, _, err := crypto.GenerateKeyPair(
		crypto.Ed25519, // Select your key type. Ed25519 are nice short
		-1,             // Select key length when possible (i.e. RSA).
	)
	addr := fmt.Sprintf("/ip4/%s/tcp/%s", deps.Config.Server.Host, deps.Config.Server.Port)
	var defaultNode host.Host
	defaultNode, err = libp2p.New(ctx,
		libp2p.ListenAddrStrings(addr),
		libp2p.Identity(defaultConfigPriv),
		libp2p.Ping(true),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)),
	)
	return defaultNode, err
}

func runCusomNodeConfig() {
	// TODO: understand libp2p config
	ctx := context.Background()
	var idht *dht.IpfsDHT
	manualConfigPriv, _, err := crypto.GenerateKeyPair(
		crypto.Ed25519, // Select your key type. Ed25519 are nice short
		3,              // Select key length when possible (i.e. RSA).
	)
	node, err := libp2p.New(ctx,
		// Use the keypair we generated
		libp2p.Identity(manualConfigPriv),
		// Multiple listen addresses
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/63785", "/ip4/0.0.0.0/tcp/63786/ws",
		),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support noise connections
		libp2p.Security(noise.ID, noise.New),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			idht, err = dht.New(ctx, h)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		// If you want to help other peers to figure out if they are behind
		// NATs, you can launch the server-side of AutoNAT too (AutoRelay
		// already runs the client)
		//
		// This service is highly rate-limited and should not cause any
		// performance issues.
		libp2p.EnableNATService(),
	)
	defer node.Close()
	// log.Printf("(WIP) Manual config host ID is %s\n", node.ID())
}
