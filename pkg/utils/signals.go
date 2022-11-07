package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func SetupSignalHandler() context.Context {
	// panic when called twice
	close(onlyOneSignalHandler)

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)

	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1) // second signal. Exit directly
	}()

	return ctx
}
