package action

import (
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/auth/service"
	"github.com/HotPotatoC/twitter-clone/pkg/config"
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
	accessToken, err := a.service.Execute(rt)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please login to continue",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken.String(),
		Expires:  accessToken.ExpiresAt(),
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
		Domain:   config.GetString("APP_DOMAIN", ""),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": accessToken.String(),
	})
}
