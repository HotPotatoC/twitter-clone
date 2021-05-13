package action

import (
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/auth/service"
	"github.com/gofiber/fiber/v2"
)

type logoutAction struct {
	service service.LogoutService
}

func NewLogoutAction(service service.LogoutService) module.Action {
	return logoutAction{service: service}
}

func (a logoutAction) Execute(c *fiber.Ctx) error {
	rt := c.Cookies("refresh_token")
	if rt == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please login to continue",
		})
	}

	err := a.service.Execute(rt)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please login to continue",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
