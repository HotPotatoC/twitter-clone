package action

import (
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/auth/service"
	"github.com/gofiber/fiber/v2"
)

type meAction struct {
	service service.MeService
}

func NewMeAction(service service.MeService) module.Action {
	return meAction{service: service}
}

func (a meAction) Execute(c *fiber.Ctx) error {
	user, err := a.service.Execute(c.Locals("accessToken").(string))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please try and login again",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": user.ID,
		"name":    user.Name,
		"handle":  user.Handle,
		"email":   user.Email,
	})
}
