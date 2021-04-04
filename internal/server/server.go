package server

import (
	"fmt"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/auth"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet"
	"github.com/HotPotatoC/twitter-clone/internal/module/user"
	"github.com/HotPotatoC/twitter-clone/pkg/cache"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/webserver"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
)

type Server struct {
	config    *Config
	db        database.Database
	cache     cache.Cache
	log       *zap.SugaredLogger
	webserver webserver.WebServer
}

// New creates a new instance of a fiber web server
func New(webserver webserver.WebServer, db database.Database, cache cache.Cache, log *zap.SugaredLogger, config *Config) *Server {
	config.init()
	return &Server{
		config:    config,
		db:        db,
		cache:     cache,
		log:       log,
		webserver: webserver,
	}
}

func (s *Server) Listen() {
	s.initMiddlewares()
	s.initRoutes()
	s.log.Infof("Starting up %s %s:%s", s.config.AppName, s.config.Version, s.config.BuildID)
	if err := s.webserver.Listen(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)); err != nil {
		s.log.Error(err)
	}
}

func (s *Server) ListenTLS(certFile string, keyFile string) {
	s.initMiddlewares()
	s.initRoutes()
	s.log.Infof("Starting up %s %s:%s", s.config.AppName, s.config.Version, s.config.BuildID)
	if err := s.webserver.ListenTLS(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port), certFile, keyFile); err != nil {
		s.log.Error(err)
	}
}

func (s *Server) initMiddlewares() {
	s.webserver.Engine().Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Content-Type",
	}))
	s.webserver.Engine().Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
	}))
	s.webserver.Engine().Use(logger.New())
}

func (s *Server) initRoutes() {
	tweetsGroup := s.webserver.Engine().Group("/tweets")
	authGroup := s.webserver.Engine().Group("/auth")
	usersGroup := s.webserver.Engine().Group("/users")
	relationshipGroup := s.webserver.Engine().Group("/relationships")
	auth.Routes(authGroup, s.db, s.cache)
	user.Routes(usersGroup, s.db, s.cache)
	tweet.Routes(tweetsGroup, s.db, s.cache)
	relationship.Routes(relationshipGroup, s.db, s.cache)
}
