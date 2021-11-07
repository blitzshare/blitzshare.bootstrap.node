package app

import (
	"github.com/libp2p/go-libp2p-core/host"
	log "github.com/sirupsen/logrus"

	dep "github.com/blitzshare/blitzshare.bootstrap.node/app/dependencies"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/p2p"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/services"
)

func Start(deps *dep.Dependencies) host.Host {
	node, err := p2p.RunNode(deps)
	if err != nil {
		panic(err)
	}
	log.Printf("(WORKING) host")
	// go PrintState(deafaultNode)
	log.Printf(" - NODE ID: %s", node.ID())
	log.Printf(" - NODE ADDR: %v", node.Addrs())
	event := services.NewNodeJoinedEvent(string(node.ID()))
	msgId, err := services.EmitNodeJoinedEvent(deps.Config.Settings.QueueUrl, event)

	if err != nil {
		log.Errorln(err)
		log.Printf("FAILED to emit node joined msg")
	} else {
		log.Printf("node joined msgId: %s", msgId)
	}
	return node
}
