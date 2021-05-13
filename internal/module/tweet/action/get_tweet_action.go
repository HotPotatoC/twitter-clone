package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type getTweetAction struct {
	service service.GetTweetService
}

func NewGetTweetAction(service service.GetTweetService) module.Action {
	return getTweetAction{service: service}
}

func (a getTweetAction) Execute(c *fiber.Ctx) error {
	tweetID, err := c.ParamsInt("tweetID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	userID := c.Locals("userID").(float64)

	tweet, err := a.service.Execute(int64(userID), int64(tweetID))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrTweetDoesNotExist):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Tweet not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(tweet)
}
