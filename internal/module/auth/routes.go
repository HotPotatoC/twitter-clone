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
	r.Post("/login", makeLoginHandler(db))
	r.Get("/me", authMiddleware.Execute(), makeMeHandler(db))
	r.Get("/token", makeTokenHandler(db, cache))
	r.Post("/logout", makeLogoutHandler(cache))
}

func makeLoginHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewLoginService(db)
		action := action.NewLoginAction(service)

		return action.Execute(c)
	}
}

func makeMeHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewMeService(db)
		action := action.NewMeAction(service)

		return action.Execute(c)
	}
}

func makeTokenHandler(db database.Database, cache cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewTokenService(db, cache)
		action := action.NewTokenAction(service)

		return action.Execute(c)
	}
}

func makeLogoutHandler(cache cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewLogoutService(cache)
		action := action.NewLogoutAction(service)

		return action.Execute(c)
	}
}
