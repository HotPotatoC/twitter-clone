package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/modules/user/entity"
	"github.com/HotPotatoC/twitter-clone/internal/modules/user/service"
	"github.com/gofiber/fiber/v2"
)

type registerAction struct {
	service service.RegisterService
}

func NewRegisterAction(service service.RegisterService) registerAction {
	return registerAction{
		service: service,
	}
}

func (a registerAction) Execute(c *fiber.Ctx) error {
	var input service.RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := input.Validate(); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	err := a.service.Execute(input)
	if err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "User with that email already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully registered a new account",
	})
}
