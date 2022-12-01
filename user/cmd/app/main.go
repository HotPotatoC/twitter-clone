package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/HotPotatoC/twitter-clone/user/clients"
	"github.com/HotPotatoC/twitter-clone/user/config"
	"github.com/HotPotatoC/twitter-clone/user/internal/server"
	"github.com/HotPotatoC/twitter-clone/user/internal/service"
	"github.com/HotPotatoC/twitter-clone/user/logger"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger.Init(true)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Wait for kill signals to gracefully shutdown the server
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-c
		cancel()
	}()

	cfg := config.New()

	clients, err := clients.NewClients(ctx, cfg)
	if err != nil {
		logger.M.Fatal(err.Error())
	}

	service := service.NewService(clients)
	server := server.New(cfg, service)

	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		logger.M.Info("starting user server")
		return server.ListenAndServe()
	})

	// Cleanups on shutdown
	group.Go(func() error {
		<-groupCtx.Done()
		logger.M.Warn("Shutting down server")
		return server.Shutdown(context.Background())
	})

	if err := group.Wait(); err != nil {
		logger.M.Warnf("Exit reason: %v\n", err)
	}
}
