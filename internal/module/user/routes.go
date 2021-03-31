package user

import (
	"github.com/HotPotatoC/twitter-clone/internal/module/user/action"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/HotPotatoC/twitter-clone/pkg/cache"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, db database.Database, cache cache.Cache) {
	r.Post("/register", buildRegisterHandler(db))
	r.Get("/:userID", buildGetUserHandler(db))
}

func buildRegisterHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewRegisterService(db)
		action := action.NewRegisterAction(service)

		return action.Execute(c)
	}
}

func buildGetUserHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewGetUserService(db)
		action := action.NewGetUserAction(service)

		return action.Execute(c)
	}
}
