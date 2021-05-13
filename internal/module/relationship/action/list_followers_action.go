package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship/service"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/gofiber/fiber/v2"
)

type listFollowersAction struct {
	service service.ListFollowersService
}

func NewListFollowersAction(service service.ListFollowersService) module.Action {
	return listFollowersAction{service: service}
}

func (a listFollowersAction) Execute(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	followers, err := a.service.Execute(int64(userID))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserDoesNotExist):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_items": len(followers),
		"items":       followers,
	})
}
