package action

import (
	"errors"
	"strconv"

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
	userID, err := strconv.ParseInt(c.Params("userID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	followers, err := a.service.Execute(userID)
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
