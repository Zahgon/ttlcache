package main

import (
	"context"
	"os/signal"
	"syscall"
)

func main() {
	shutdown := runServer()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	<-ctx.Done()

	shutdown()
}

// runServer starts the HTTP server and returns a shutdown function.
func runServer() func() { _ = "STUB: not implemented"; return nil }
