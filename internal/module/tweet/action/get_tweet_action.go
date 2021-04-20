package action

import (
	"errors"
	"fmt"
	"strconv"

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
	tweetID, err := strconv.ParseInt(c.Params("tweetID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}
	tweet, err := a.service.Execute(tweetID)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, entity.ErrTweetDoesNotExist) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Tweet not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(tweet)
}
