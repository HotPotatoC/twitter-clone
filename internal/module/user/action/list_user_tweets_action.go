package action

import (
	"errors"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/gofiber/fiber/v2"
)

type listUserTweetsAction struct {
	service service.ListUserTweetsService
}

func NewListUserTweetsAction(service service.ListUserTweetsService) module.Action {
	return listUserTweetsAction{service: service}
}

func (a listUserTweetsAction) Execute(c *fiber.Ctx) error {
	createdAtCursor := c.Query("cursor")
	username := c.Params("username")

	userID := c.Locals("userID").(float64)

	tweets, err := a.service.Execute(int64(userID), username, createdAtCursor)
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
		"items":       tweets,
		"total_items": len(tweets),
	})
}
