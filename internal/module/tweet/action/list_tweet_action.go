package action

import (
	"errors"

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
	createdAtCursor := c.Query("cursor")
	tweets, err := a.service.Execute(createdAtCursor)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCursor) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Invalid cursor",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(tweets),
		"items":       tweets,
	})
}
