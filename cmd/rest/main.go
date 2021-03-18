package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/HotPotatoC/twitter-clone/internal/server"
	"github.com/HotPotatoC/twitter-clone/pkg/config"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/logger"
	"github.com/HotPotatoC/twitter-clone/pkg/version"
	"github.com/HotPotatoC/twitter-clone/pkg/webserver"
	"github.com/gofiber/fiber/v2"
)

var (
	debug   bool
	host    string
	port    string
	cfgPath string
	prefork bool
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Run the app as DEBUG mode")
	flag.BoolVar(&debug, "d", false, "Run the app as DEBUG mode")

	flag.StringVar(&host, "host", "127.0.0.1", "The http server's host")
	flag.StringVar(&host, "h", "127.0.0.1", "The http server's host")

	flag.StringVar(&port, "port", "5000", "The http server's port")
	flag.StringVar(&port, "p", "5000", "The http server's port")

	flag.StringVar(&cfgPath, "config", "./configs/.env", "The application configurations")
	flag.StringVar(&cfgPath, "c", "./configs/.env", "The application configurations")

	flag.BoolVar(&prefork, "prefork", false, "Run the app in Prefork mode [multiple Go processes]")
}

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfgPath, err := filepath.Abs(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	config.Load(cfgPath)

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.GetString("DB_USER", "postgres"),
		config.GetString("DB_PASSWORD", ""),
		config.GetString("DB_HOST", "127.0.0.1"),
		config.GetInt("DB_PORT", 5432),
		config.GetString("DB_DATABASE", "twitterclone"),
	)

	db, err := database.New(ctx, dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger(config.GetBool("DEBUG", false))
	webserver := webserver.New(fiber.Config{
		Prefork: prefork,
	})

	server := server.New(webserver, db, logger, &server.Config{
		AppName: config.GetString("APP_NAME", "twitter-clone"),
		Version: version.Version,
		Host:    host,
		Port:    port,
	})

	server.Listen(fmt.Sprintf("%s:%s", host, port))
}
