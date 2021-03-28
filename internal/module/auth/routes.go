package auth

import (
	"github.com/HotPotatoC/twitter-clone/internal/module/auth/action"
	"github.com/HotPotatoC/twitter-clone/internal/module/auth/service"
	"github.com/HotPotatoC/twitter-clone/internal/server/middleware"
	"github.com/HotPotatoC/twitter-clone/pkg/cache"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, db database.Database, cache cache.Cache) {
	authMiddleware := middleware.NewAuthMiddleware()
	r.Post("/login", buildLoginHandler(db))
	r.Get("/me", authMiddleware.Execute(), buildMeHandler(db))
	r.Get("/token", buildTokenHandler(db))
}

func buildLoginHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewLoginService(db)
		action := action.NewLoginAction(service)

		return action.Execute(c)
	}
}

func buildMeHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewMeService(db)
		action := action.NewMeAction(service)

		return action.Execute(c)
	}
}

func buildTokenHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewTokenService(db)
		action := action.NewTokenAction(service)

		return action.Execute(c)
	}
}
