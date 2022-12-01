package server

import (
	"net/http"

	"github.com/HotPotatoC/twitter-clone/user/config"
	"github.com/HotPotatoC/twitter-clone/user/internal/service"
	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// New creates a new user twirp server
func New(cfg *config.Config, service service.Service) http.Server {
	handler := newHandler(service)
	userServiceServer := user.NewUserServiceServer(handler)

	mux := chi.NewMux()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	mux.Mount(userServiceServer.PathPrefix(), userServiceServer)

	return http.Server{
		Addr:    cfg.App.Address,
		Handler: mux,
	}
}

type Handler interface {
	user.UserService
}

type handler struct {
	service service.Service
}

func newHandler(service service.Service) Handler {
	return &handler{service: service}
}
