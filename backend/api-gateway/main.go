package main

import (
	"api-gateway/pkg/di"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := di.InitDI()
	server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.Shutdown()
}
