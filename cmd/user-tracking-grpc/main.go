package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RyuChk/ips-map-service/cmd/user-tracking-grpc/di"
	"github.com/RyuChk/ips-map-service/internal/config"
)

func main() {
	config.LoadConfig()

	appName := "user-tracking-service"
	if err := os.Setenv("APP_NAME", appName); err != nil {
		log.Panic("error while setting APP_NAME")
	}

	_, cleanUpFunc, err := di.InitializeContainer()
	if err != nil {
		log.Panic("Error While building user-tracking-service grpc server")
	}
	defer cleanUpFunc()

	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, os.Interrupt, syscall.SIGTERM)

	<-closeCh
}
