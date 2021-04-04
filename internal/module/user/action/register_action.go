package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/HotPotatoC/twitter-clone/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type registerAction struct {
	service service.RegisterService
}

func NewRegisterAction(service service.RegisterService) module.Action {
	return registerAction{service: service}
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

	accessToken, refreshToken, err := a.service.Execute(input)
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

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.String(),
		Expires:  refreshToken.ExpiresAt(),
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
		Domain:   config.GetString("APP_DOMAIN", ""),
		SameSite: "None",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Successfully registered a new account",
		"access_token": accessToken.String(),
	})
}
