package server

import (
	"fmt"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/aws"
	"github.com/HotPotatoC/twitter-clone/internal/common/cache"
	"github.com/HotPotatoC/twitter-clone/internal/common/config"
	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/common/webserver"
	"github.com/HotPotatoC/twitter-clone/internal/module/auth"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet"
	"github.com/HotPotatoC/twitter-clone/internal/module/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"go.uber.org/zap"
)

type Server struct {
	config    *Config
	s3        *aws.S3Bucket
	db        database.Database
	cache     cache.Cache
	log       *zap.SugaredLogger
	webserver webserver.WebServer
}

// New creates a new instance of a fiber web server
func New(webserver webserver.WebServer, s3 *aws.S3Bucket, db database.Database, cache cache.Cache, log *zap.SugaredLogger, config *Config) *Server {
	config.init()
	return &Server{
		config:    config,
		db:        db,
		s3:        s3,
		cache:     cache,
		log:       log,
		webserver: webserver,
	}
}

func (s *Server) Listen() {
	s.initMiddlewares()
	s.initRouteGroups()
	if !fiber.IsChild() {
		s.log.Infof("Starting up %s", s.config.AppName)
	}
	if err := s.webserver.Listen(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)); err != nil {
		s.log.Error(err)
	}
}

func (s *Server) ListenTLS(certFile string, keyFile string) {
	s.initMiddlewares()
	s.initRouteGroups()
	if !fiber.IsChild() {
		s.log.Infof("Starting up %s", s.config.AppName)
	}
	if err := s.webserver.ListenTLS(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port), certFile, keyFile); err != nil {
		s.log.Error(err)
	}
}

func (s *Server) initMiddlewares() {
	s.webserver.Engine().Use(cors.New(cors.Config{
		AllowOrigins:     config.GetString("APP_DOMAIN", "*"),
		AllowCredentials: true,
		AllowHeaders:     "Content-Type",
	}))

	s.webserver.Engine().Use(helmet.New())

	s.webserver.Engine().Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	s.webserver.Engine().Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
	}))

	s.webserver.Engine().Use(logger.New(logger.Config{
		Format: "${green}${time}${reset} | ${status} | ${cyan}${latency}${reset}	-	${host} | ${yellow}${method}${reset} | ${path} ${queryParams}\n",
	}))
}

func (s *Server) initRouteGroups() {
	auth.Routes(
		s.webserver.Engine().Group("/auth"),
		s.db,
		s.cache)

	tweet.Routes(
		s.webserver.Engine().Group("/tweets"),
		s.db,
		s.s3,
		s.cache)

	user.Routes(
		s.webserver.Engine().Group("/users"),
		s.db,
		s.s3,
		s.cache)

	relationship.Routes(
		s.webserver.Engine().Group("/relationships"),
		s.db,
		s.cache)
}
