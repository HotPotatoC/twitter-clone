package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/service"
	"github.com/gofiber/fiber/v2"
)

type favoriteTweetAction struct {
	service service.FavoriteTweetService
}

func NewFavoriteTweetAction(service service.FavoriteTweetService) module.Action {
	return favoriteTweetAction{service: service}
}

func (a favoriteTweetAction) Execute(c *fiber.Ctx) error {
	tweetID, err := c.ParamsInt("tweetID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
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
		case errors.Is(err, entity.ErrTweetAlreadyFavorited):
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"message": "Successfully unfavorited a tweet",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully favorited a tweet",
	})
}
