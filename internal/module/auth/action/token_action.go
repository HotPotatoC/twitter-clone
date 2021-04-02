package action

import (
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/auth/service"
	"github.com/gofiber/fiber/v2"
)

type tokenAction struct {
	service service.TokenService
}

func NewTokenAction(service service.TokenService) module.Action {
	return tokenAction{service: service}
}

func (a tokenAction) Execute(c *fiber.Ctx) error {
	rt := c.Cookies("refresh_token")
	if rt == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please login to continue",
		})
	}

	accessToken, err := a.service.Execute(rt)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please login to continue",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": accessToken.String(),
	})
}
