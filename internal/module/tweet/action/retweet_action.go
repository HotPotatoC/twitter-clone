package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type retweetAction struct {
	service service.RetweetService
}

func NewRetweetAction(service service.RetweetService) module.Action {
	return retweetAction{service: service}
}

func (a retweetAction) Execute(c *fiber.Ctx) error {
	tweetID, err := c.ParamsInt("tweetID")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}
	userID := c.Locals("userID").(float64)

	err = a.service.Execute(int64(tweetID), int64(userID))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrTweetDoesNotExist):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Tweet not found",
			})
		case errors.Is(err, entity.ErrTweetAlreadyRetweeted):
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"message": "Successfully removed retweet",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully retweeted a tweet",
	})
}
