package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type listTweetFeedAction struct {
	service service.ListTweetFeedService
}

func NewListTweetFeedAction(service service.ListTweetFeedService) module.Action {
	return listTweetFeedAction{service: service}
}

func (a listTweetFeedAction) Execute(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)
	createdAtCursor := c.Query("cursor")
	tweets, err := a.service.Execute(int64(userID), createdAtCursor)
	if err != nil {
		switch {
		case errors.Is(err, module.ErrInvalidCursor):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Invalid cursor",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(tweets),
		"items":       tweets,
	})
}
