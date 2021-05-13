package action

import (
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type listTweetRepliesAction struct {
	service service.ListTweetRepliesService
}

func NewListTweetRepliesAction(service service.ListTweetRepliesService) module.Action {
	return listTweetRepliesAction{service: service}
}

func (a listTweetRepliesAction) Execute(c *fiber.Ctx) error {
	tweetID, err := c.ParamsInt("tweetID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	createdAtCursor := c.Query("cursor")
	replies, err := a.service.Execute(int64(tweetID), createdAtCursor)
	if err != nil {
		fmt.Println(err)
		switch {
		case errors.Is(err, module.ErrInvalidCursor):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Invalid cursor",
			})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(replies),
		"items":       replies,
	})
}
