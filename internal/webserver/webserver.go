package webserver

import (
	"github.com/gofiber/fiber/v2"
)

type WebServer interface {
	Listen(addr string) error
	ListenTLS(addr string, certFile string, keyFile string) error
	Shutdown() error
	Engine() *fiber.App
}

type webserver struct {
	engine *fiber.App
}

// New creates a new instance of a fiber web server
func New(config ...fiber.Config) WebServer {
	return &webserver{engine: fiber.New(config...)}
}

// Listen starts the webserver
func (s *webserver) Listen(addr string) error {
	return s.engine.Listen(addr)
}

// ListenTLS starts the webserver in https
func (s *webserver) ListenTLS(addr string, certFile string, keyFile string) error {
	return s.engine.ListenTLS(addr, certFile, keyFile)
}

// Shutdown stop the server
func (s *webserver) Shutdown() error {
	return s.engine.Shutdown()
}

func (s *webserver) Engine() *fiber.App {
	return s.engine
}
