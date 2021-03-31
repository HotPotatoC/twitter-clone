package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/HotPotatoC/twitter-clone/internal/server"
	"github.com/HotPotatoC/twitter-clone/pkg/cache"
	"github.com/HotPotatoC/twitter-clone/pkg/config"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/logger"
	"github.com/HotPotatoC/twitter-clone/pkg/version"
	"github.com/HotPotatoC/twitter-clone/pkg/webserver"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var (
	cfgPath  string
	prefork  bool
	https    bool
	certFile string
	keyFile  string
)

func init() {
	flag.StringVar(&cfgPath, "config", "./configs/.env", "The application configurations")
	flag.StringVar(&cfgPath, "c", "./configs/.env", "The application configurations")

	flag.BoolVar(&prefork, "prefork", false, "Run the app in Prefork mode [multiple Go processes]")
	flag.BoolVar(&https, "https", false, "Run the app in HTTPS mode")
	flag.StringVar(&certFile, "certFile", "./configs/server.crt", "Public key file path")
	flag.StringVar(&keyFile, "keyFile", "./configs/server.key", "Private key file path")
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

	cache := cache.NewRedisClient(ctx, &redis.Options{
		Addr:     config.GetString("REDIS_ADDR", "localhost:6379"),
		Password: config.GetString("REDIS_PASSWORD", ""),
	})

	if err := cache.Ping(); err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger(config.GetBool("DEBUG", false))
	webserver := webserver.New(fiber.Config{
		Prefork: prefork,
	})

	server := server.New(webserver, db, cache, logger, &server.Config{
		Version: version.Version,
		BuildID: version.BuildID,
		AppName: config.GetString("APP_NAME", "twitter-clone"),
		Host:    config.GetString("APP_HOST", "127.0.0.1"),
		Port:    config.GetString("APP_PORT", "5000"),
	})

	if https {
		server.ListenTLS(certFile, keyFile)
	} else {
		server.Listen()
	}
}
