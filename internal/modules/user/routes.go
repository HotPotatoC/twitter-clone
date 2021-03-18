package user

import (
	"github.com/HotPotatoC/twitter-clone/internal/modules/user/action"
	"github.com/HotPotatoC/twitter-clone/internal/modules/user/service"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, db database.Database) {
	r.Post("/register", buildRegisterHandler(db))
}

func buildRegisterHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewRegisterService(db)
		action := action.NewRegisterAction(service)

		return action.Execute(c)
	}
}
