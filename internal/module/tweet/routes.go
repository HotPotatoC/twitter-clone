package tweet

import (
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/action"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/HotPotatoC/twitter-clone/internal/server/middleware"
	"github.com/HotPotatoC/twitter-clone/pkg/cache"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, db database.Database, cache cache.Cache) {
	authMiddleware := middleware.NewAuthMiddleware()
	r.Post("/", authMiddleware.Execute(), buildCreateTweetHandler(db))
}

func buildCreateTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewCreateTweetService(db)
		action := action.NewCreateTweetAction(service)

		return action.Execute(c)
	}
}
