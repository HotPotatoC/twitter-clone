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
	r.Get("/", buildListTweetHandler(db))
	r.Post("/", authMiddleware.Execute(), buildCreateTweetHandler(db))
	r.Get("/:tweetID/replies", buildListTweetRepliesHandler(db))
	r.Post("/:tweetID/reply", authMiddleware.Execute(), buildCreateReplyHandler(db))
	r.Post("/:tweetID/favorite", authMiddleware.Execute(), buildFavoriteTweetHandler(db))
}

func buildListTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewListTweetService(db)
		action := action.NewListTweetAction(service)

		return action.Execute(c)
	}
}

func buildCreateTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewCreateTweetService(db)
		action := action.NewCreateTweetAction(service)

		return action.Execute(c)
	}
}

func buildListTweetRepliesHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewListTweetRepliesService(db)
		action := action.NewListTweetRepliesAction(service)

		return action.Execute(c)
	}
}

func buildCreateReplyHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewCreateReplyService(db)
		action := action.NewCreateReplyAction(service)

		return action.Execute(c)
	}
}

func buildFavoriteTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewFavoriteTweetService(db)
		action := action.NewFavoriteTweetAction(service)

		return action.Execute(c)
	}
}
