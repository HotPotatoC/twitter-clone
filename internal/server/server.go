package server

import (
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/modules/user"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/webserver"
	"go.uber.org/zap"
)

type Server struct {
	config    *Config
	db        database.Database
	log       *zap.SugaredLogger
	webserver webserver.WebServer
}

// New creates a new instance of a fiber web server
func New(webserver webserver.WebServer, db database.Database, log *zap.SugaredLogger, config *Config) *Server {
	config.init()
	return &Server{
		config:    config,
		db:        db,
		log:       log,
		webserver: webserver,
	}
}

func (s *Server) Listen() {
	s.initRoutes()
	s.log.Infof("Starting up %s %s:%s", s.config.AppName, s.config.Version, s.config.BuildID)
	if err := s.webserver.Listen(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)); err != nil {
		s.log.Error(err)
	}
}

func (s *Server) initRoutes() {
	tweetsGroup := s.webserver.Engine().Group("/tweets")
	authGroup := s.webserver.Engine().Group("/auth")
	usersGroup := s.webserver.Engine().Group("/users")
	followersGroup := s.webserver.Engine().Group("/followers")
	user.Routes(usersGroup, s.db)
	_ = tweetsGroup
	_ = authGroup
	_ = followersGroup
}
