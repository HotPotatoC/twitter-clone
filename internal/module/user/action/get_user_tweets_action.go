package action

import (
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/gofiber/fiber/v2"
)

type getUserTweetsAction struct {
	service service.GetUserTweetsService
}

func NewGetUserTweetsAction(service service.GetUserTweetsService) module.Action {
	return getUserTweetsAction{service: service}
}

func (a getUserTweetsAction) Execute(c *fiber.Ctx) error {
	createdAtCursor := c.Query("cursor")
	username := c.Params("username")

	userID := c.Locals("userID").(float64)

	tweets, err := a.service.Execute(int64(userID), username, createdAtCursor)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, entity.ErrUserDoesNotExist) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There was a problem on our side",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"items":       tweets,
		"total_items": len(tweets),
	})
}
