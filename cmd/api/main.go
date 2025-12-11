package main

import (
	"log"
	"os"
	"os/signal"
	"project/internal/infra/config"
	"syscall"

	_ "project/docs"
)

// @title API
// @version 0.0.1
// @description API
// @contact.name Ruham Leal
// @contact.email ruhamxlpro@hotmail.com
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8085
// @BasePath /
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
