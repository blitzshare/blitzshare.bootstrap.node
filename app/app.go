package app

import (
	dep "github.com/blitzshare/blitzshare.bootstrap.node/app/dependencies"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/p2p"
)

func Start(d *dep.Dependencies) {
	p2p.RunNode(d)
}
