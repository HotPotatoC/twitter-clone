package action

import (
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/gofiber/fiber/v2"
)

type updateUserAction struct {
	service service.UpdateUserService
}

func NewUpdateUserAction(service service.UpdateUserService) module.Action {
	return updateUserAction{service: service}
}

func (a updateUserAction) Execute(c *fiber.Ctx) error {
	var input service.UpdateUserInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := input.Validate(); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	userID := c.Locals("userID").(float64)
	err := a.service.Execute(input, int64(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"messages": "Successfully updated user data",
	})
}
