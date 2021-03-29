package action

import (
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type listTweetAction struct {
	service service.ListTweetService
}

func NewListTweetAction(service service.ListTweetService) module.Action {
	return listTweetAction{service: service}
}

func (a listTweetAction) Execute(c *fiber.Ctx) error {
	tweets, err := a.service.Execute()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(tweets),
		"items":       tweets,
	})
}
