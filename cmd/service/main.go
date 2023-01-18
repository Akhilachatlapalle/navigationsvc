package main

import (
	// stdlib
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	// third-party
	"go.uber.org/zap"

	// local
	"github.com/akhilachatlapalle/navigationsvc/internal/config"
	"github.com/akhilachatlapalle/navigationsvc/internal/service"
)

func main() {
	flag.Parse()

	conf, err := config.New()
	if err != nil {
		log.Fatalln("Failed to instantiate configuration:", err)
	}

	defer func() { _ = conf.Logger.Sync() }()
	conf.Logger.Info("Initializing service")

	srv, err := service.NewServer(context.Background(), conf)
	if err != nil {
		conf.Logger.Fatal("Failed to initialize service", zap.Error(err))
	}

	go func() {
		sig := AwaitSignals()
		conf.Logger.Info("Received signal", zap.String("signal", sig.String()))

		conf.Logger.Info("Stopping service")
		if err := srv.Stop(); err != nil {
			conf.Logger.Error("Failed to stop service", zap.Error(err))
		}
		conf.Logger.Info("Stopped service")
	}()

	conf.Logger.Info("Starting service", zap.String("server_url", srv.URL()))

	if err := srv.Start(); err != nil { // Blocking call.
		conf.Logger.Error("Error while running service", zap.Error(err))
	}
}

func AwaitSignals() os.Signal {
	ch := make(chan os.Signal, 2)
	return awaitSignals(ch)
}

func awaitSignals(ch chan (os.Signal)) os.Signal {
	signal.Notify(ch, syscall.SIGTERM, os.Interrupt)
	return <-ch
}
