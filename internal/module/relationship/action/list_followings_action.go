package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship/service"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/gofiber/fiber/v2"
)

type listFollowingsAction struct {
	service service.ListFollowingsService
}

func NewListFollowingsAction(service service.ListFollowingsService) module.Action {
	return listFollowingsAction{service: service}
}

func (a listFollowingsAction) Execute(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	followings, err := a.service.Execute(int64(userID))
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
		"total_items": len(followings),
		"items":       followings,
	})
}
