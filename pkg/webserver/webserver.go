package webserver

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type WebServer interface {
	Listen(addr string) error
	Engine() *fiber.App
}

type webserver struct {
	engine *fiber.App
}

// New creates a new instance of a fiber web server
func New(config ...fiber.Config) WebServer {
	return &webserver{
		engine: fiber.New(config...),
	}
}

// Listen starts the webserver
func (s *webserver) Listen(addr string) error {
	go func(addr string) error {
		if err := s.engine.Listen(addr); err != nil {
			return err
		}
		return nil
	}(addr)

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	if err := s.engine.Shutdown(); err != nil {
		return err
	}
	return nil
}

func (s *webserver) Engine() *fiber.App {
	return s.engine
}
