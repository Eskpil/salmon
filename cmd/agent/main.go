package main

import (
	"time"

	"github.com/eskpil/salmon/cmd/agent/mycontext"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, err := mycontext.NewContext()

	if err != nil {
		log.Fatalf("Failed to create a new context: %v\n", err)
	}

	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})

	log.Info("Performing initial routine")
	go ctx.PerformRoutine()

	log.Info("Starting main loop")
	for {
		select {
		case <-ticker.C:
			go ctx.PerformRoutine()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
