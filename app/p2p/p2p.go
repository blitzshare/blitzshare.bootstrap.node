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
