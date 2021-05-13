package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/gofiber/fiber/v2"
)

type getUserAction struct {
	service service.GetUserService
}

func NewGetUserAction(service service.GetUserService) module.Action {
	return getUserAction{service: service}
}

func (a getUserAction) Execute(c *fiber.Ctx) error {
	username := c.Params("username")
	userID := c.Locals("userID").(float64)

	user, err := a.service.Execute(int64(userID), username)
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

	return c.Status(fiber.StatusOK).JSON(user)
}
