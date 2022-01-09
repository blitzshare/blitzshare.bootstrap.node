package app

import (
	"strconv"

	"github.com/libp2p/go-libp2p-core/host"
	log "github.com/sirupsen/logrus"

	dep "github.com/blitzshare/blitzshare.bootstrap.node/app/dependencies"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/p2p"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/services"
)

func Start(deps *dep.Dependencies) host.Host {
	log.Printf("staring node")
	node, err := p2p.RunNode(deps)

	if err != nil {
		log.Errorln("Shwopsy")
		panic(err)
	}
	log.Printf("(WORKING) host")
	// go PrintState(deafaultNode)
	log.Printf(" - NODE ID: %s", node.ID())
	log.Printf(" - NODE ADDR: %v", node.Addrs())
	port, _ := strconv.Atoi(deps.Config.Server.Port)
	event := services.NewNodeRegistryCmd(node.ID().Pretty(), port)
	msgId, err := services.EmitNodeRegistryCmd(deps.Config.Settings.QueueUrl, event)
	if err != nil {
		log.Errorln(err)
		log.Printf("FAILED to emit event")
	} else {
		log.Printf("Node joined msgId: %s", msgId)
	}
	return node
}
