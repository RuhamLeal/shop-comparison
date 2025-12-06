package main

import (
	"log"
	"os"
	"os/signal"
	"project/internal/infra/config"
	"syscall"
)

func main() {
	environmentConf := config.NewBaseConfig(".env")
	mainInstance := config.NewServerInstances(environmentConf)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		log.Print("Shutting down server...")
		mainInstance.Stop()
	}()

	mainInstance.Start()
}
