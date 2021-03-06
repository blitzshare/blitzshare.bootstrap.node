package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/blitzshare/blitzshare.bootstrap.node/app"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/config"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/dependencies"
	log "github.com/sirupsen/logrus"
)

func initLog() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// log.SetLevel(log.WarnLevel)
}

func main() {
	log.Printf("running main")
	cfg, err := config.Load()
	initLog()
	if err != nil {
		log.Fatalf("failed to load config %v\n", err)
	}
	deps, err := dependencies.NewDependencies(cfg)
	if err != nil {
		log.Fatalf("failed to load dependencies %v\n", err)
	}
	node := app.Start(deps)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	// shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}

}
