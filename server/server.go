package server

import (
	"fmt"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/aws"
	"github.com/HotPotatoC/twitter-clone/internal/cache"
	"github.com/HotPotatoC/twitter-clone/internal/config"
	"github.com/HotPotatoC/twitter-clone/internal/database"
	"github.com/HotPotatoC/twitter-clone/internal/webserver"
	"github.com/HotPotatoC/twitter-clone/module/auth"
	"github.com/HotPotatoC/twitter-clone/module/relationship"
	"github.com/HotPotatoC/twitter-clone/module/tweet"
	"github.com/HotPotatoC/twitter-clone/module/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	s.initRoutes()
	if !fiber.IsChild() {
		s.log.Infof("Starting up %s %s:%s", s.config.AppName, s.config.Version, s.config.BuildID)
	}
	if err := s.webserver.Listen(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)); err != nil {
		s.log.Error(err)
	}
}

func (s *Server) ListenTLS(certFile string, keyFile string) {
	s.initMiddlewares()
	s.initRoutes()
	if !fiber.IsChild() {
		s.log.Infof("Starting up %s %s:%s", s.config.AppName, s.config.Version, s.config.BuildID)
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

	s.webserver.Engine().Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
	}))

	s.webserver.Engine().Use(logger.New(logger.Config{
		Format: "${green}${time}${reset} | ${status} | ${cyan}${latency}${reset}	-	${host} | ${yellow}${method}${reset} | ${path} ${queryParams}\n",
	}))
}

func (s *Server) initRoutes() {
	authGroup := s.webserver.Engine().Group("/auth")
	tweetsGroup := s.webserver.Engine().Group("/tweets")
	usersGroup := s.webserver.Engine().Group("/users")
	relationshipGroup := s.webserver.Engine().Group("/relationships")
	auth.Routes(authGroup, s.db, s.cache)
	tweet.Routes(tweetsGroup, s.db, s.cache)
	user.Routes(usersGroup, s.db, s.s3, s.cache)
	relationship.Routes(relationshipGroup, s.db, s.cache)
}
